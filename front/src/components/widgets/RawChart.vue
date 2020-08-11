<template>
  <v-row>
    <v-col cols="12" sm="12" md="8">
      <canvas id="ei_raw_data"></canvas>
      <raw-expenses-table v-if="selected_rows.length > 0" :rows="selected_rows"></raw-expenses-table>
    </v-col>
    <v-col class="cat_list_wrapper" cols="12" sm="12" md="4">
      <div class="cat_list" v-for="cat in cats">
        <v-checkbox class="cat_list_item shrink mr-2"
                    input-value="true"
                    :label="cat.title"
                    @change="updateConditions(cat.id)"
                    v-model="chart_config[cat.id.toString()]"></v-checkbox>
        <div class="cat_sub_list" v-for="sCat in cat.sub">
          <v-checkbox class="cat_list_item"
                      input-value="false"
                      :label="sCat.title"
                      @change="updateConditions(sCat.id)"
                      v-model="chart_config[sCat.id.toString()]"></v-checkbox>
        </div>
      </div>
    </v-col>
  </v-row>
</template>

<script>
import RawExpensesTable from "./RawExpensesTable";
export default {
  name: "RawChart",
  components: {RawExpensesTable},
  data() {
    return {
      raw_rows: [],
      parsed_row: [],
      chart_config: {},
      table_data: [],
      chart: null,
      selected_rows: [],
    }
  },
  computed: {
    cats() {
      return this.$store.state.categories_expenses || [];
    },
    cats_inc() {
      return this.$store.state.categories_incoming || [];
    },
  },
  created() {
    this.getSelectedConfig();
    this.loadData();
  },
  methods: {
    updateConditions(id) {
      let targetData = this.cats.find(v => {
        if (v.id === id) {
          return true;
        }
        for (let i of v.sub) {
          if (i.id === id) {
            return true;
          }
        }
        return false;
      });
      if (!targetData) {
        return false;
      }

      if (targetData.id === id) {
        let subVal = this.chart_config[id.toString()];
        for (let j of targetData.sub) {
          this.chart_config[j.id.toString()] = !subVal;
        }
      } else {
        this.chart_config[targetData.id.toString()] = false;
      }

      this.$store.commit('setRawStatisticConfig', this.chart_config);
    },
    getSelectedConfig() {
      let currentConfig = this.$store.getters.statisticRawConfig || {};
      if (Object.keys(currentConfig).length !== 0) {
        console.log('load from storage');
        this.chart_config = currentConfig;
        return;
      }

      for (let i of this.cats) {
        this.chart_config[i.id] = true;
        for (let j of i.sub) {
          this.chart_config[j.id] = true;
        }
      }
      this.$store.commit('setRawStatisticConfig', this.chart_config);
      console.log('Set selected all', this.chart_config, this.cats);
    },

    generateDataSet(rows) {
      let tmp = {};
      let monthList = [];
      // prepare draft of amount
      for (let i of rows) {
        let cur = tmp[i.category] || [];
        let month = this.$moment(+i.created_at * 1000).format('YYYY-MM');
        let amount = cur[month] || 0;
        cur[month] = amount + i.amount;
        tmp[i.category] = cur;
        monthList.push(month);
      }
      monthList = [...new Set(monthList)];
      let colorIndex = 1;
      let dataSet = [];

      for (let i of this.cats) {
        if (!this.chart_config[i.id] && i.sub.length === 0) {
          // it is disabled root category without any child
          continue;
        } else if (this.chart_config[i.id]) {
          // count root category with child
          let color = (`color_${colorIndex}` in window.chartColors) ? colorIndex : 1;
          ++colorIndex;
          let obj = {
            label: i.title,
            //type: 'line',
            stack: 'line 1',
            backgroundColor: window.chartColors[`color_${color}`],
            data: this.countRootWithSubCats(i, tmp, monthList),
          };
          dataSet.push(obj);
        } else {
          // each child is separated on graph
        }
      }

      let result = {};
      for (let month of monthList) {
        result[month] = 0;
      }
      for (let i of this.cats_inc) {
        let amountData = tmp[i.id] || [];
        for (let month of monthList) {
          result[month] += amountData[month] || 0;
        }
      }
      let resp = [];
      for (let i of Object.keys(result)) {
        resp.push(result[i]);
      }

      dataSet.push({
        label: 'Доходы',
        type: 'line',
        stack: 'line 2',
        backgroundColor: window.chartColors.incoming_color,
        data: resp,
      });

      return {data: dataSet, month: monthList};
    },

    countRootWithSubCats(rootCategory, tmp, monthList) {
      let amountData = tmp[rootCategory.id] || [];
      let result = {};
      for (let month of monthList) {
        result[month] = amountData[month] || 0;
      }
      for (let i of rootCategory.sub) {
        amountData = tmp[i.id] || [];
        for (let month of monthList) {
          result[month] += amountData[month] || 0;
        }
      }
      let resp = [];
      for (let i of Object.keys(result)) {
        resp.push(result[i]);
      }
      return resp;
    },

    initChart(rows) {

      let dataSet = this.generateDataSet(rows)
      let ctx = document.getElementById('ei_raw_data').getContext('2d');
      this.chart = new Chart(ctx, {
        type: 'bar',
        data: {
          labels: dataSet.month,
          datasets: dataSet.data,
        },
        options: {
          legend: {
            display: false,
            position: 'right',
          },
          title: {
            display: true,
            text: this.$t('raw_visualization')
          },
          onClick: this.loadCategoryData,
        },
        scales: {
          xAxes: [{
            stacked: true,
          }],
          yAxes: [{
            stacked: true,
          }]
        }
      });
    },

    loadCategoryData(e, i) {
      if (!this.chart) {
        return;
      }
      let activeElement = this.chart.getElementAtEvent(e);
      let selectedLabel = activeElement[0]._view.datasetLabel;
      let selectedMonth = activeElement[0]._view.label;

      let dateSort = this.$moment(`${selectedMonth}-01`);
      let startTimestamp = dateSort.unix();
      let endTimestamp = dateSort.endOf('month').unix();

      for (let i of this.cats) {
        if (i.title === selectedLabel) {
          // it is root category
          let ids = {};
          ids[i.id] = i.title;
          for (let j of i.sub) {
            ids[j.id] = j.title;
          }
          let rows = [];
          let catIds = Object.keys(ids);
          for (let j = 0; j < this.raw_rows.length; j++) {
            if (!catIds.includes(this.raw_rows[j].category.toString())) {
              continue;
            }
            let validDate = this.raw_rows[j].created_at >= startTimestamp && this.raw_rows[j].created_at <= endTimestamp
            if (!validDate) {
              continue;
            }
            let tmp = this.raw_rows[j];
            tmp.amount_expense = '-' + tmp.amount;
            tmp.amount_incoming = '';
            tmp.cat_name = ids[tmp.category];
            tmp.created_at = this.$moment(+tmp.created_at * 1000).format('YYYY-MM-DD HH:mm');
            rows.push(tmp);
          }
          this.selected_rows = rows;
        }
        for (let j of i.sub) {
          if (j.title === selectedLabel) {
            // it is child category
          }
        }
      }
    },

    loadData() {
      this.askBackend('data/statistics/money_change', {}).then(
          resp => {
            if (!resp.ok) {
              return
            }
            this.raw_rows = resp.rows;
            this.initChart(resp.rows);
          }
      )
    }
  }
}
</script>

<style lang="scss">
  .cat_list_wrapper {
    max-width: 500px;
  }
  .cat_list {
    font-size: 14px;
    .cat_sub_list {
      margin-left: 30px;
    }
  }
  .cat_list_item {
    margin: 0;
    .v-input__control {
      .v-input__slot {
        margin-bottom: 0;
        i, label {
          font-size: 14px;
        }
      }
      .v-messages {
        display: none;
      }
    }
  }
  .v-input__slot {
    margin-bottom: 0;
  }
</style>
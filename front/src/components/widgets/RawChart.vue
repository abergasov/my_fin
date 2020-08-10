<template>
  <v-row>
    <v-col cols="12" sm="12" md="8">
      <canvas id="ei_raw_data"></canvas>
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
export default {
  name: "RawChart",
  data() {
    return {
      parsed_row: [],
      chart_config: {},
      table_data: [],
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
        console.log(amountData);
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
      new Chart(ctx, {
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
          }
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

    loadData() {
      this.askBackend('data/statistics/money_change', {}).then(
          resp => {
            if (!resp.ok) {
              return
            }
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
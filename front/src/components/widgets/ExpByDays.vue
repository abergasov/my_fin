<template>
  <v-row no-gutters>
    <v-col cols="12" sm="12" md="8">
      <div class="chart_wrapper">
        <canvas id="exp_by_day"></canvas>
      </div>
    </v-col>
    <v-col class="dd_table" cols="12" sm="12" md="4">
      <v-simple-table>
        <template v-slot:default>
          <tbody>
          <tr v-for="item in dataDays" :key="item.x">
            <td>{{ item.x }}</td>
            <td>{{ item.y }}</td>
          </tr>
          <tr>
            <td>{{ $t('total') }}</td>
            <td>{{ total }}</td>
          </tr>
          </tbody>
        </template>
      </v-simple-table>
    </v-col>
  </v-row>
</template>

<script>
  export default {
    name: "ExpByDays",
    data () {
      return {
        isSmallScreen: false,
        dataDays: [],
        total: 0,
      }
    },
    created() {
      window.addEventListener("main_page_load", this.loadData);
    },
    beforeDestroy() {
      window.removeEventListener("main_page_load", this.loadData);
    },
    methods: {
      initChart(data) {
        new Chart(document.getElementById('exp_by_day').getContext('2d'), {
          type: 'bar',
          data: {
            datasets: [{
              data: data,
              backgroundColor: 'rgba(183, 28, 28, 1)',
              borderColor: [
                'rgba(76, 175, 80, 1)',
                'rgba(183, 28, 28, 1)',
                'rgba(255, 159, 64, 1)',

              ],
            }],

          },
          options: {
            maintainAspectRatio: false,
            scales: {
              xAxes: [{
                type: 'time',
                time: {
                  unit: 'day'
                }
              }]
            },
            legend: {
              display: false,
            },

            title: {
              display: true,
              text: this.$t('last_14_days')
            }
          }
        });
      },

      loadData() {
        console.log('event listener exp by days')
        this.dataDays = this.$store.state.per_days_chart;
        this.total = this.dataDays.map(i => i.y).reduce((prev, next) => prev + next);
        this.initChart(this.dataDays);
      }
    }
  }
</script>

<style scoped lang="scss">
  .chart_wrapper {
    height: 330px;
  }
  .dd_table {
    .v-data-table {
      table {
        td {
          height: 20px;
        }
      }
    }
  }
</style>
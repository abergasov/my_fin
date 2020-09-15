<template>
  <v-row no-gutters>
    <v-col cols="12" sm="12" md="8">
      <div class="chart_wrapper">
        <canvas id="exp_by_day"></canvas>
      </div>
    </v-col>
    <v-col cols="12" sm="12" md="4">
      awdawd
    </v-col>
  </v-row>
</template>

<script>
  export default {
    name: "ExpByDays",
    data () {
      return {
        isSmallScreen: false
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
        // data = JSON.stringify(data);
        // data = JSON.parse(data);
        console.log(data);
        console.log(data);
        console.log(data);
        console.log(data);
        console.log(data);
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
        let resp = this.$store.state.per_days_chart;
        this.initChart(resp);
      }
    }
  }
</script>

<style scoped lang="css">
  .chart_wrapper {
    height: 300px;
  }
</style>
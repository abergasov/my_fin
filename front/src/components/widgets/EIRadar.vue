<template>
    <canvas id="ei_radar"></canvas>
</template>

<script>
    export default {
        name: "EIRadar",
        created() {
            window.addEventListener("main_page_load", this.loadData);
        },
        beforeDestroy() {
          window.removeEventListener("main_page_load", this.loadData);
        },
      methods: {
            initChart(data, percent, percent_optional) {
              let ctx = document.getElementById('ei_radar').getContext('2d');
              new Chart(ctx, {
                  type: 'pie',
                  data: {
                      datasets: [{
                          data: data,
                          backgroundColor: [
                              'rgba(76, 175, 80, 0.2)',
                              'rgba(183, 28, 28, 0.2)',
                              'rgba(255, 159, 64, 0.2)',
                          ],
                          borderColor: [
                              'rgba(76, 175, 80, 1)',
                              'rgba(183, 28, 28, 1)',
                              'rgba(255, 159, 64, 1)',

                          ],
                      }],

                      labels: [
                          this.$t('graph_label_incoming') + ' ' + data[0],
                          this.$t('graph_label_expense') + ' ' + percent + '%',
                          this.$t('graph_label_expense_mandatory') + ' ' + percent_optional + '%',
                      ]
                  },
                  options: {
                      legend: {
                          position: 'right',
                      },
                      title: {
                          display: true,
                          text: this.$t('last_30_days')
                      }
                  }
              });
          },

            loadData() {
                console.log('event listener')
                let resp = this.$store.state.el_chart;
                this.initChart(resp.rows, resp.percent, resp.percent_optional);
            }
        }
    }
</script>

<style scoped>

</style>
<template>
    <canvas id="myChart"></canvas>
</template>

<script>
    export default {
        name: "EIRadar",
        created() {
           // setTimeout(() => {
                this.loadData();
           // }, 1000);
        },
        methods: {
            initChart(data) {
              let ctx = document.getElementById('myChart').getContext('2d');
              new Chart(ctx, {
                  type: 'pie',
                  data: {
                      datasets: [{
                          data: data,
                          backgroundColor: [
                              'rgba(76, 175, 80, 0.2)',
                              'rgba(183, 28, 28, 0.2)',
                              'rgba(54, 162, 235, 0.2)',
                              'rgba(255, 206, 86, 0.2)',
                              'rgba(75, 192, 192, 0.2)',
                              'rgba(153, 102, 255, 0.2)',
                              'rgba(255, 159, 64, 0.2)'
                          ],
                          borderColor: [
                              'rgba(76, 175, 80, 1)',
                              'rgba(183, 28, 28, 1)',
                              'rgba(54, 162, 235, 1)',
                              'rgba(255, 206, 86, 1)',
                              'rgba(75, 192, 192, 1)',
                              'rgba(153, 102, 255, 1)',
                              'rgba(255, 159, 64, 1)'
                          ],
                      }],

                      labels: [
                          this.$t('graph_label_incoming'),
                          this.$t('graph_label_expense'),
                      ]
                  },
                  options: {}
              });
          },

            loadData() {
                this.askBackend('data/statistics/list', {}).then(
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

<style scoped>

</style>
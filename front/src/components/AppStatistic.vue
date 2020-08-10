<template>
  <v-card id="statistic_container" class="mx-auto" outlined>
    <raw-chart></raw-chart>
  </v-card>
</template>

<script>
import RawChart from "./widgets/RawChart";
export default {
  name: "AppStatistic",
  components: {RawChart},
}
</script>

<style lang="scss">
  #statistic_container {
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
  }
</style>
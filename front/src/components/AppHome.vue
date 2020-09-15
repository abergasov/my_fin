<template>
  <v-container>
    <v-row no-gutters>
      <v-col cols="12" sm="12" md="4" lg="3">
        <money-change></money-change>
      </v-col>
      <v-col cols="12" sm="12" md="4" lg="3">
        <EIRadar></EIRadar>
      </v-col>
      <v-col cols="12" sm="12" md="4" lg="3">
        <router-link to="/">Home</router-link> |
        <router-link to="/medium">Medium</router-link>
        <router-view />
      </v-col>
      <v-col cols="12" sm="12" md="4" lg="3">
        <DebtList></DebtList>
      </v-col>
    </v-row>
    <ExpByDays></ExpByDays>
    <ExpensesTable></ExpensesTable>
  </v-container>
</template>

<script>
  import ExpensesTable from "./widgets/ExpensesTable";
  import MoneyChange from "./widgets/MoneyChange";
  import EIRadar from "./widgets/EIRadar";
  import ExpByDays from "./widgets/ExpByDays";
  import DebtList from "./widgets/DebtList";
  export default {
    components: {DebtList, ExpensesTable, MoneyChange, EIRadar, ExpByDays},
    name: 'AppHome',
    data () {
      return {
        dialog: false,
        incoming: false,
        expense: false,
        asset: false,
        category: null,
        amount: '',
        commentary: '',
      }
    },
    created() {
      this.loadAll();
    },
    methods: {

      loadAll() {
        this.askBackend('data/bulk/main', {})
            .then(data => {
              if (data.ok) {
                this.$store.commit('setExpenses', data.expenses || []);
                this.$store.commit('setElChart', data.ei_radar || {});
                this.$store.commit('setPerDays', data.per_day || []);
                window.dispatchEvent(new Event('main_page_load'));
              }
            })
      },
    }
  };
</script>

<style scoped lang="scss">
    h1  {
        color: green;
    }

</style>
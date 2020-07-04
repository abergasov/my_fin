<template>
  <v-container>
    <v-row no-gutters>
      <v-col cols="12" sm="12" md="4" lg="3">
        <money-change></money-change>
      </v-col>
      <v-col cols="12" sm="12" md="4" lg="3">
        <router-link to="/">Home</router-link> |
        <router-link to="/medium">Medium</router-link>
        <router-view />
      </v-col>
    </v-row>
    <ExpensesTable></ExpensesTable>
  </v-container>
</template>

<script>
  import ExpensesTable from "./widgets/ExpensesTable";
  import MoneyChange from "./widgets/MoneyChange";
  export default {
    components: {ExpensesTable, MoneyChange},
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
    methods: {







    }
  };
</script>

<style scoped lang="scss">
    h1  {
        color: green;
    }

</style>
<template>
  <v-row no-gutters class="table_wrapper">
    <v-card outlined>
      <v-card-title class="headline">
        {{ $t('changes') }}
        <v-spacer></v-spacer>
        <v-text-field
            v-model="search"
            append-icon="mdi-magnify"
            :label="$t('search')"
            single-line
            hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table :headers="headers"
                    ref="tableEl"
                    :items="rows"
                    :search="search"
                    :sort-by="['calories', 'fat']"
                    :sort-desc="[false, true]"
                    multi-sort
                    :items-per-page="(rows.length > 30 ? 30 : -1)"
                    class="elevation-1">
        <template slot="body.append">
          <tr class="pink--text">
            <th class="title">Totals</th>
            <th class="title">{{ total_ex_sum }}</th>
          </tr>
        </template>
      </v-data-table>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: "RawExpensesTable",
  props: {
    rows: Array,
  },
  data () {
    return {
      search: '',
      headers: [
        { text: 'Category', value: 'cat_name' },
        { text: 'Expense', value: 'amount_expense' },
        { text: 'Created at', value: 'created_att' },
        { text: 'Commentary', value: 'commentary' },
      ],
      total_ex_sum: '',
    }
  },
  mounted() {
    this.calculateTotal(1500);
  },
  watch: {
    search() {
      this.calculateTotal();
    }
  },
  methods: {
    calculateTotal(timeout) {
      let time = timeout || 400;
      setTimeout(() => {
        this.total_ex_sum = this.sumField('amount_expense');
      }, time)
    },
    sumField(key) {
      if (!this.$refs.tableEl) {
        return
      }
      if (!this.$refs.tableEl.$children.length > 0) {
        return
      }
      if (!this.$refs.tableEl.$children[0].computedItems.length > 0) {
        return
      }

      return this.$refs.tableEl.$children[0].computedItems.reduce((accumulator, currentValue) => {
        if (currentValue[key].length > 0) {
          return accumulator + Math.abs(currentValue[key]);
        }
        return accumulator;
      }, 0);
    },
  }
}
</script>

<style scoped lang="scss">
.table_wrapper {
  margin-top: 15px;
  .v-card {
    width: 100%;
    .active_row {
    }
    .expense_row {
    }
    .incoming_row {
    }
  }
}
</style>
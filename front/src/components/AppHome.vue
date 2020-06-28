<template>
  <v-container>
    <v-row no-gutters>
      <v-col cols="12" sm="12" md="4" lg="3">
        <v-card outlined>
          <v-card-title class="headline">{{ $t('capital_change') }}</v-card-title>
          <v-card-actions>
            <ul class="actions_list">
              <li><v-btn @click="addIncome" x-large color="success" block>{{ $t('add_income') }}</v-btn></li>
              <li><v-btn @click="addExpense" x-large color="error" block>{{ $t('add_expense') }}</v-btn></li>
            </ul>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col cols="12" sm="12" md="4" lg="3">

      </v-col>
      <div>
        <router-link to="/">Home</router-link> |
        <router-link to="/medium">Medium</router-link>
        <router-view />
      </div>
      <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
          <v-toolbar :color="getClassAction()">
            <v-btn icon @click="closeAdd">
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>{{ $t(getTitleAction()) }}</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
              <v-btn :disabled="formValid()" text @click="saveAdd">{{ $t('capital_change_add') }}</v-btn>
            </v-toolbar-items>
          </v-toolbar>
          <v-card class="mx-auto data_insert" max-width="600" outlined>
            <div class="headline mb-1">{{ $t('select_category') }}</div>
            <v-divider></v-divider>
            <v-radio-group class="cat_select" v-model="category">
              <div class="root_cat" v-for="c in categories" :key="c.id">
                <v-radio class="gender" :label="c.title" :value="c.id"></v-radio>
                <div class="sub_cat" v-for="s in c.sub" :key="s.id">
                  <v-radio :label="s.title" :value="s.id"></v-radio>
                </div>
              </div>
            </v-radio-group>
            <v-divider></v-divider>
            <v-text-field type="number" v-model="amount" :label="$t('amount')" class="amount_input" hide-details="auto"></v-text-field>
            <v-divider></v-divider>
            <v-textarea filled :label="$t('commentary')" rows="2"></v-textarea>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn @click="saveAdd" :color="(incoming ? 'success' : 'error')" :disabled="formValid()">
                <v-icon>mdi-contrast</v-icon>{{ $t('add') }}</v-btn>
            </v-card-actions>
          </v-card>
        </v-card>
      </v-dialog>
    </v-row>
    <ExpensesTable></ExpensesTable>
  </v-container>
</template>

<script>
  import ExpensesTable from "./widgets/ExpensesTable";
  export default {
    components: {ExpensesTable},
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

    },
    computed: {
      categories() {
        return this.$store.state.categories;
      }
    },
    methods: {
      formValid() {
        return !(+this.amount > 0 && this.category !== null);
      },
      addIncome() {
        this.dialog = true;
        this.incoming = true;
      },
      addExpense() {
        this.dialog = true;
        this.expense = true;
      },
      addAsset() {
        this.dialog = true;
        this.asset = true;
      },
      closeAdd() {
        this.dialog = false;
        this.incoming = false;
        this.expense = false;
        this.asset = false;
        this.category = null;
        this.amount = '';
        this.commentary = '';
      },

      getClassAction() {
        switch (this.getType()) {
          case 'E':
            return 'error';
          case 'A':
            return 'info';
          case 'I':
            return 'success';
        }
      },

      getTitleAction() {
        switch (this.getType()) {
          case 'E':
            return 'add_expense';
          case 'A':
            return 'add_active';
          case 'I':
            return 'add_income';
        }
      },

      getType() {
        let type = 'E';
        if (this.incoming) {
          type = 'I';
        } else if (this.asset) {
          type = 'A';
        }
        return type;
      },

      saveAdd() {
        this.askBackend('data/expense/add', {
          cat: +this.category,
          amount: +this.amount,
          incoming: this.getType(),
          commentary: this.commentary,
        }).then(({data}) => {
          this.$store.commit('setAlert', {
            display: true,
            text: (data.ok ?  this.$t('added') : this.$t('not_added')),
            color: (data.ok ? 'success' : 'error'),
            delay: 2,
          });
        });
        this.closeAdd();
      }
    }
  };
</script>

<style scoped lang="scss">
    h1  {
        color: green;
    }
    ul.actions_list {
      list-style: none;
      width: 100%;
      padding-left: 0;
      li:not(:first-child) {
        margin-top: 10px;
      }
    }

    div.sub_cat {
      margin-left: 30px;
      margin-top: 5px;
      margin-bottom: 5px;
      div.v-radio {
        label.v-label {
          font-size: 25px;
        }
      }

    }

    div.root_cat {
      margin-bottom: 5px;
      div.v-radio {
        label.v-label {
          font-size: 25px;
        }
      }
    }

    .data_insert {
      padding: 15px;
      hr.v-divider {
        margin-top: 10px;
        margin-bottom: 10px;
      }
    }

    .cat_select {
      margin-top: 0;
      margin-left: 30px;
    }

    div.amount_input {
      .v-label {
        padding-left: 10px;
      }

    }
</style>
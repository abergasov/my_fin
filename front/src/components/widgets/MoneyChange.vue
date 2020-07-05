<template>
    <div>
        <v-card outlined>
            <v-card-actions>
                <ul class="actions_list">
                    <li><v-btn @click="addIncome" x-large color="success" block>{{ $t('add_income') }}</v-btn></li>
                    <li><v-btn @click="addExpense" x-large color="error" block>{{ $t('add_expense') }}</v-btn></li>
                    <li><AssetChange></AssetChange></li>
                </ul>
            </v-card-actions>
        </v-card>
        <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
            <v-card>
                <v-toolbar :color="this.incoming ? 'success' : 'error'">
                    <v-btn icon @click="closeAdd">
                        <v-icon>mdi-close</v-icon>
                    </v-btn>
                    <v-toolbar-title>{{ $t((this.incoming ? 'add_income' : 'add_expense')) }}</v-toolbar-title>
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
    </div>
</template>

<script>
    import AssetChange from './AssetChange'
    export default {
        components: {
            AssetChange
        },
        name: "MoneyChange",
        data () {
            return {
                dialog: false,
                dialog_asset: false,
                incoming: false,
                active: false,
                expense: false,
                category: null,
                amount: '',
                commentary: '',
            }
        },
        computed: {
            asset_category() {
                return this.$store.state.categories_assets;
            },
            categories() {
                return this.incoming ? this.$store.state.categories_incoming : this.$store.state.categories_expenses;
            },
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
                this.dialog_asset = true;
                this.active = true;
            },
            closeAdd() {
                this.dialog_asset = false;
                this.dialog = false;
                this.incoming = false;
                this.expense = false;
                this.active = false;
                this.category = null;
                this.amount = '';
                this.commentary = '';
            },
            saveAdd() {
                this.askBackend('data/expense/add', {
                    cat: +this.category,
                    amount: +this.amount,
                    incoming: this.getType(+this.category),
                    commentary: this.commentary,
                }).then(data => {
                    this.$store.commit('setAlert', {
                        display: true,
                        text: (data.ok ?  this.$t('added') : this.$t('not_added')),
                        color: (data.ok ? 'success' : 'error'),
                        delay: 2,
                    });
                });
                this.closeAdd();
            },
            getType(catId) {
                if (this.incoming) {
                    return 'I';
                }
                let type = 'E';
                for (let i = 0; i < this.categories.length; i++) {
                    let po = this.categories[i];
                    if (po.id === catId) {
                        type = po.cat_type === 'Em' ? 'Em' : 'E';
                        break;
                    }
                    for (let j = 0; j < po.sub.length; j++) {
                        let o = po.sub[j];
                        if (o.id === catId) {
                            type = o.cat_type === 'Em' ? 'Em' : 'E';
                            break;
                        }
                    }
                }
                return type;
            }
        }
    }
</script>

<style scoped lang="scss">
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
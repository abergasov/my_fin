<template>
    <v-card outlined>
        <v-card-actions>
            <ul class="actions_list">
                <li><v-btn class="ma-2" @click="openDialog('add')" outlined color="success">{{ $t('add_debt') }}</v-btn></li>
                <li><v-btn class="ma-2" @click="openDialog('give')" outlined color="error">{{ $t('give_debt') }}</v-btn></li>
            </ul>
            <v-spacer></v-spacer>
            <v-chip class="ma-2" @click="debt_list = true" outlined color="teal" text-color="white">
                <v-avatar v-if="countActive() > 0" left class="green darken-4">{{ countActive() }}</v-avatar>
                {{ $t('debts') }}
            </v-chip>
        </v-card-actions>
        <v-dialog v-model="debt_list"  max-width="800px">
            <v-card>
                <v-card-title>
                    <span class="headline">{{ $t('debt_list') }}</span>
                    <v-spacer></v-spacer>
                    <v-text-field
                            v-model="search"
                            append-icon="mdi-magnify"
                            :label="$t('search')"
                            single-line
                            hide-details
                    ></v-text-field>
                </v-card-title>
                <v-card-text>
                    <v-data-table :headers="debt_headers"
                                  :items="debts"
                                  :search="search"
                                  :loading="isLoading"
                    >
                        <template v-slot:top>
                            <v-dialog v-if="edited_debt_item" v-model="debt_edit_dialog" max-width="500px">
                                <v-card>
                                    <v-card-title>
                                        <span class="headline">{{ $t('change_debt_status') }}</span>
                                    </v-card-title>
                                    <v-card-text>
                                        <v-container>
                                            <v-row>
                                                <v-col cols="4">{{ edited_debt_item.amount }}</v-col>
                                                <v-col cols="8">
                                                    <p class="font-weight-regular">{{ edited_debt_item.commentary }}</p>
                                                </v-col>
                                            </v-row>
                                            <v-row>
                                                <v-col cols="12" sm="12" md="12">
                                                    <v-switch v-model="edited_debt_item.active_debt" inset :label="$t('debt_payed')"></v-switch>
                                                </v-col>
                                            </v-row>
                                        </v-container>
                                    </v-card-text>

                                    <v-card-actions>
                                        <v-spacer></v-spacer>
                                        <v-btn color="blue darken-1" text @click="closeDebtItem">Cancel</v-btn>
                                        <v-btn color="blue darken-1" text @click="saveDebtItem">Save</v-btn>
                                    </v-card-actions>
                                </v-card>
                            </v-dialog>
                        </template>
                        <template v-slot:item.amount="{ item }">
                            <v-chip v-if="item.debt_type !== 1 && +item.active_debt !== 1" color="error" class="debt_chip" dark>
                                {{ item.amount }}
                            </v-chip>
                            <span v-else>{{ item.amount }}</span>
                        </template>
                        <template v-slot:item.active_debt="{ item }">
                            <v-simple-checkbox v-model="+item.active_debt !== 1" disabled></v-simple-checkbox>
                        </template>
                        <template v-slot:item.edit_it="{ item }">
                            <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
                        </template>
                    </v-data-table>
                </v-card-text>
            </v-card>
        </v-dialog>
        <v-dialog v-model="debt_dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
            <v-card>
                <v-toolbar :color="this.incoming ? 'success' : 'error'">
                    <v-btn icon @click="closeAdd">
                        <v-icon>mdi-close</v-icon>
                    </v-btn>
                    <v-toolbar-title>{{ $t((this.incoming ? 'add_debt' : 'give_debt')) }}</v-toolbar-title>
                    <v-spacer></v-spacer>
                    <v-toolbar-items>
                        <v-btn :disabled="formValid()" text @click="saveAdd">{{ $t('capital_change_add') }}</v-btn>
                    </v-toolbar-items>
                </v-toolbar>
                <v-card class="mx-auto data_insert" max-width="600" outlined>
                    <div class="headline mb-1">{{ $t('debt_date') }}</div>
                    <v-divider></v-divider>
                    <v-date-picker full-width v-model="picker" :color="this.incoming ? 'success' : 'error'"></v-date-picker>
                    <v-divider></v-divider>
                    <v-text-field type="number" v-model="amount" :label="$t('amount')" class="amount_input" hide-details="auto"></v-text-field>
                    <v-divider></v-divider>
                    <v-textarea v-model="commentary" filled :label="$t('commentary')" rows="2"></v-textarea>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn @click="saveAdd" :color="(incoming ? 'success' : 'error')" :disabled="formValid()">
                            <v-icon>mdi-contrast</v-icon>{{ $t('add') }}</v-btn>
                    </v-card-actions>
                </v-card>
            </v-card>
        </v-dialog>
    </v-card>
</template>

<script>
    export default {
        name: "DebtList",
        data() {
            return {
                search: '',
                debt_headers: [
                    { text: this.$t('amount'), value: 'amount' },
                    { text: this.$t('type'), value: 'type_h' },
                    { text: this.$t('created_at'), value: 'created_at_h' },
                    { text: this.$t('commentary'), value: 'commentary' },
                    { text: this.$t('debt_date'), value: 'payment_date_h' },
                    { text: this.$t('active_debt'), value: 'active_debt' },
                    { text: '', value: 'edit_it' },
                ],
                debt_edit_dialog: false,
                edited_debt_item: null,
                debt_dialog: false,
                debt_list: false,
                incoming: false,
                amount: '',
                commentary: '',
                picker: '',
            }
        },
        created() {
            this.askBackend('data/debt/get', {}).then(
                resp => {
                    this.$store.commit('setDebts', resp.debts);
                }
            );
        },
        computed: {
            debts() {
                let rows = this.$store.state.debts;
                for (let i = 0; i < rows.length; i++) {
                    let tmp = rows[i];
                    rows[i].type_h = +tmp.debt_type !== 1 ? this.$t('debt_type_credit') : this.$t('debt_type_debt');
                    rows[i].created_at_h = this.$moment(+tmp.created_at * 1000).format('YYYY-MM-DD');
                    rows[i].payment_date_h = this.$moment(+tmp.payment_date * 1000).format('YYYY-MM-DD');
                }
                return rows;
            },
            isLoading() {
                return this.$store.state.dataLoading;
            },
        },
        methods: {
            countActive() {
                let c = 0;
                for (let i in this.debts) {
                    if (!this.debts.hasOwnProperty(i)) {
                        continue;
                    }
                    if (this.debts[i].active_debt !== 1) {
                        c += 1;
                    }
                }
                return c;
            },
            editItem(item) {
                this.edited_debt_item = {
                    commentary: item.commentary,
                    active_debt: +item.active_debt !== 1,
                    amount: +item.amount,
                    debt_id: item.debt_id,
                };
                this.debt_edit_dialog = true;
            },
            closeDebtItem() {
                this.edited_debt_item = null;
                this.debt_edit_dialog = false;
            },
            saveDebtItem() {
                this.askBackend('data/debt/pay', {
                    'debt_id': this.edited_debt_item.debt_id,
                    'debt_active': (this.edited_debt_item.active_debt ? 0 : 1),
                }).then(
                    resp => {
                        if (resp.ok) {
                            this.$store.commit('setDebts', resp.debts);
                        }
                    }
                )

                this.closeDebtItem();
            },

            openDialog(type) {
                this.incoming = type !== 'give';
                this.debt_dialog = true;
                this.amount = '';
                this.commentary = '';
                this.picker = '';
            },
            closeAdd() {
                this.debt_dialog = false;
            },
            formValid() {
                return !(+this.amount > 0 && this.commentary.length > 0 && this.picker.length > 0);
            },
            saveAdd() {
                this.askBackend('data/debt/add', {
                    amount: +this.amount,
                    debt_type: this.incoming ? 1 : 0,
                    commentary: this.commentary,
                    payment_date: this.$moment(this.picker).unix(),
                }).then(data => {
                    if (data.ok) {
                        this.$store.commit('setDebts', data.debts);
                    }
                    this.$store.commit('setAlert', {
                        display: true,
                        text: (data.ok ?  this.$t('added') : this.$t('not_added')),
                        color: (data.ok ? 'success' : 'error'),
                        delay: 5,
                    });
                });
                this.closeAdd();
            }
        }
    }
</script>

<style scoped lang="scss">
    ul.actions_list {
        list-style: none;
        li {
            button.v-btn {
                min-width: 146px;
            }
        }
    }
    .amount_input {
        margin-top: 20px;
        margin-bottom: 20px;
        padding-bottom: 12px;
    }
    .debt_chip {
        min-width: 63px;
    }
</style>
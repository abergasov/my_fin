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
                          :loading="isLoading"
                          :sort-by="['calories', 'fat']"
                          :sort-desc="[false, true]"
                          multi-sort
                          class="elevation-1">
                <template slot="body.append">
                    <tr class="pink--text">
                        <th class="title">Totals</th>
                        <th class="title">{{ total_ex_sum }}</th>
                        <th class="title">{{ total_in_sum }}</th>
                    </tr>
                </template>
            </v-data-table>
        </v-card>
    </v-row>
</template>

<script>
    export default {
        name: "ExpensesTable",
        data () {
            return {
                search: '',
                headers: [
                    { text: 'Category', value: 'cat_name' },
                    { text: 'Expense', value: 'amount_expense' },
                    { text: 'Incoming', value: 'amount_incoming' },
                    { text: 'Type', value: 'type' },
                    { text: 'Created at', value: 'created_at' },
                    { text: 'Commentary', value: 'commentary' },
                ],
                total_ex_sum: '',
                total_in_sum: '',
            }
        },
        mounted() {
            this.calculateTotal(1500);
        },
        computed: {
            rows() {
                let rows = this.$store.state.expenses;
                let mixed = [];
                let exCat = this.simplyCat(this.cats);
                let inCat = this.simplyCat(this.cats_inc);
                let simplyCat = Object.assign(exCat, inCat);
                for (let i = 0; i < rows.length; i++) {
                    let tmp = rows[i];
                    tmp.amount_expense = (tmp.incoming === 'E' || tmp.incoming === 'Em') ? '-' + tmp.amount : '';
                    tmp.amount_incoming = (tmp.incoming !== 'E' && tmp.incoming !== 'Em') ? '+' + tmp.amount : '';
                    tmp.cat_name = simplyCat[rows[i].cat];
                    tmp.type = this.getType(tmp.incoming);
                    tmp.created_at = this.$moment(+tmp.created_at * 1000).format('YYYY-MM-DD HH:mm');
                    mixed.push(tmp);
                }
                return mixed;
            },
            isLoading() {
                return this.$store.state.dataLoading;
            },
            cats() {
                return this.$store.state.categories_expenses || [];
            },
            cats_inc() {
                return this.$store.state.categories_incoming || [];
            },
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
                    this.total_in_sum = this.sumField('amount_incoming');
                }, time)
            },
            simplyCat(ct) {
                let rebuild = {};
                for (let i of ct) {
                    rebuild[i.id] = i.title;
                    for (let j of i.sub) {
                        rebuild[j.id] = j.title;
                    }
                }
                return rebuild;
            },
            getType(type) {
                switch (type) {
                    case 'E':
                        return this.$t('graph_label_expense');
                    case 'Em':
                        return this.$t('graph_label_expense_mandatory');
                    default:
                        return this.$t('table_label_incoming');

                }
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

            getRowClassForType(type) {
                switch (type) {
                    case 'E'://expenses
                        return 'expense_row';
                    case 'I'://incoming
                        return 'incoming_row';
                    case 'A'://active
                        return 'active_row';
                    default:
                        return '';
                }
            }
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
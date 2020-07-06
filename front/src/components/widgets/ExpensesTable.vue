<template>
    <v-row no-gutters class="table_wrapper">
        <v-card outlined>
            <v-card-title class="headline">{{ $t('changes') }}</v-card-title>
            <v-data-table :headers="headers"
                          :items="rows"
                          :loading="isLoading"
                          :sort-by="['calories', 'fat']"
                          :sort-desc="[false, true]"
                          multi-sort
                          class="elevation-1">
            </v-data-table>
        </v-card>
    </v-row>
</template>

<script>
    import moment from "moment";
    export default {
        name: "ExpensesTable",
        data () {
            return {
                headers: [
                    { text: 'Category', value: 'cat_name' },
                    { text: 'Expense', value: 'amount_expense' },
                    { text: 'Incoming', value: 'amount_incoming' },
                    { text: 'Created at', value: 'created_at' },
                    { text: 'Commentary', value: 'commentary' },
                ],
            }
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
                    tmp.amount_expense = tmp.incoming === 'E' ? '-' + tmp.amount : '';
                    tmp.amount_incoming = tmp.incoming !== 'E' ? '+' + tmp.amount : '';
                    tmp.cat_name = simplyCat[rows[i].cat];
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
        methods: {
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
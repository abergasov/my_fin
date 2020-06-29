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
                <template v-slot:item="props">
                    <tr :class="getRowClassForType(props.item.incoming)">
                        <td>{{props.item.cat_name}}</td>
                        <td>{{ (props.item.incoming === 'E' ? '-' : '+') + props.item.amount}}</td>
                        <td>{{props.item.created_at}}</td>
                        <td style="max-width: 150px">{{props.item.commentary}}</td>
                    </tr>
                </template>
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
                    { text: 'Amount', value: 'amount' },
                    { text: 'Created at', value: 'created_at' },
                    { text: 'Commentary', value: 'commentary' },
                ],
            }
        },
        computed: {
            rows() {
                let rows = this.$store.state.expenses;
                let mixed = [];
                let simplyCat = this.simplyCat();
                for (let i = 0; i < rows.length; i++) {
                    let tmp = rows[i];
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
                return this.$store.state.categories || [];
            },
        },
        methods: {
            simplyCat() {
                let rebuild = {};
                for (let i of this.cats) {
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
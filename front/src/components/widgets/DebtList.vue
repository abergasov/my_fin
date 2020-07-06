<template>
    <v-card outlined>
        <v-card-actions>
            <ul class="actions_list">
                <li><v-btn class="ma-2" @click="openDialog('add')" outlined color="success">{{ $t('add_debt') }}</v-btn></li>
                <li><v-btn class="ma-2" @click="openDialog('give')" outlined color="error">{{ $t('give_debt') }}</v-btn></li>
            </ul>
        </v-card-actions>
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
                debt_dialog: false,
                incoming: false,
                amount: '',
                commentary: '',
                picker: '',
            }
        },
        methods: {
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
</style>
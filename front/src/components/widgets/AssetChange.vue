<template>
    <div>
    <v-btn @click="dialog_asset = !dialog_asset" x-large color="info" block>{{ $t('add_active') }}</v-btn>
    <v-dialog v-model="dialog_asset" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
            <v-toolbar color="info">
                <v-btn icon @click="closeAdd">
                    <v-icon>mdi-close</v-icon>
                </v-btn>
                <v-toolbar-title>{{ $t('add_active') }}</v-toolbar-title>
                <v-spacer></v-spacer>
                <v-toolbar-items>
                    <v-btn :disabled="formValid()" text @click="addAsset">{{ $t('capital_change_add') }}</v-btn>
                </v-toolbar-items>
            </v-toolbar>
            <v-card class="mx-auto data_insert" max-width="600" outlined>
                <div class="headline mb-1">{{ $t('select_category') }}</div>
                <v-divider></v-divider>
                <v-radio-group class="cat_select" v-model="category">
                    <div class="root_cat" v-for="c in asset_category" :key="c.id">
                        <v-radio class="gender" :label="$t(c.title)" :value="c.id"></v-radio>
                    </div>
                </v-radio-group>
                <v-divider></v-divider>
                <v-text-field type="number" v-model="amount" :label="$t('amount')" class="amount_input" hide-details="auto"></v-text-field>
                <v-divider></v-divider>
                <v-textarea filled :label="$t('commentary')" rows="2"></v-textarea>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn @click="addAsset" color="info" :disabled="formValid()">
                        <v-icon>mdi-contrast</v-icon>{{ $t('add') }}</v-btn>
                </v-card-actions>
            </v-card>
        </v-card>
    </v-dialog>
    </div>
</template>

<script>
    export default {
        name: "AssetChange",
        data () {
            return {
                dialog_asset: false,
                active: false,
                category: null,
                amount: '',
                commentary: '',
            }
        },
        computed: {
            asset_category() {
                return this.$store.state.categories_assets;
            },
        },
        methods: {
            addAsset() {

            },
            formValid() {
                return !(+this.amount > 0 && this.category !== null);
            },
            closeAdd() {
                this.dialog_asset = false;
                this.active = false;
                this.category = null;
                this.amount = '';
                this.commentary = '';
            },
        },
    }
</script>

<style scoped>
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
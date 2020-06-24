<template>
    <v-form width="600">
        <v-card class="mx-auto" width="600">
            <v-toolbar flat>
                <v-toolbar-title class="font-weight-light">{{ $t('category') }}</v-toolbar-title>
                <v-spacer></v-spacer>
                <v-btn v-if="editMode" color="error" class="cancel_top" fab small @click="cancelChanges">
                    <v-icon>mdi-close-octagon-outline</v-icon>
                </v-btn>
                <v-btn v-if="editMode" color="success" fab small @click="addRootCategory">
                    <v-icon>mdi-plus</v-icon>
                </v-btn>
                <v-btn v-else color="info" fab small @click="editMode = !editMode">
                    <v-icon>mdi-pencil</v-icon>
                </v-btn>
            </v-toolbar>
            <v-card v-if="editMode">
                <div class="editor_wrapper" v-for="c in categories_copy" :key="c.id">
                    <v-row>
                        <v-text-field v-model="c.title" :label="$t('new_category_name')" full-width hide-details="auto">
                            <v-icon slot="prepend" color="error" @click="remove(c.id)">mdi-minus</v-icon>
                            <v-icon slot="append" color="success" @click="addSubCategory(c.id)">mdi-plus</v-icon>
                        </v-text-field>
                    </v-row>
                    <v-text-field v-for="s in c.sub" :key="s.id" class="sub_draft" v-model="s.title" :label="$t('new_sub_category_name')" full-width hide-details="auto">
                        <v-icon slot="prepend" color="error" @click="remove(s.id, c.id)">mdi-minus</v-icon>
                    </v-text-field>
                </div>
            </v-card>
            <v-list v-else>
                <v-list-group v-for="c in categories_copy" no-action sub-group value="true" :key="c.id">
                    <template v-slot:activator>
                        <v-list-item-content>
                            <v-list-item-title>{{ c.title }}</v-list-item-title>
                        </v-list-item-content>
                    </template>
                    <v-list-item v-for="s in c.sub" :key="s.id">
                        <v-list-item-title v-if="!editMode">{{ s.title }}</v-list-item-title>
                    </v-list-item>
                </v-list-group>
            </v-list>
            <v-row v-if="editMode" class="button_wrapper">
                <v-btn color="error" class="mr-4" @click="cancelChanges">{{ $t('cancel') }}</v-btn>
                <v-btn color="success" class="mr-4" @click="updateCategories">{{ $t('update') }}</v-btn>
            </v-row>
        </v-card>
        <v-snackbar v-if="update_ok" :value="true" bottom color="success" outlined right>{{ $t('categories_updated') }}</v-snackbar>
        <v-snackbar v-if="update_wrong" :value="true" bottom color="error" outlined right>{{ $t('categories_updated_error') }}</v-snackbar>
    </v-form>
</template>

<script>
    export default {
        name: "AppCategories",
        data () {
            return {
                snackbar: true,
                editMode: false,
                categories_copy: [],
                update_ok: false,
                update_wrong: false,
            }
        },
        created() {
            this.getUserCategories();
        },
        methods: {
            remove(id, parentID) {
                if (parentID) {
                    let i = this.categories_copy.findIndex(c => c.id === parentID);
                    if (i === -1) {
                        return
                    }
                    let parent = this.categories_copy[i];
                    this.categories_copy[i].sub = parent.sub.filter(c => c.id !== id);
                } else {
                    this.categories_copy = this.categories_copy.filter(c => c.id !== id)
                }
            },
            addSubCategory(id) {
                let i = this.categories_copy.findIndex(c => c.id === id);
                if (i !== -1) {
                    this.categories_copy[i].sub.push(
                        this.getNewCategory()
                    );
                }
            },
            addRootCategory() {
                this.categories_copy.push(
                    this.getNewCategory()
                )
            },
            getNewCategory() {
                return {
                    title: '',
                    id: Math.floor(Math.random() * 1000),
                    sub: [],
                };
            },
            updateCategories() {
                this.askBackend('user_category/update', {cat: this.categories_copy})
                    .then(({data}) => {
                        if (data.ok) {
                            this.$store.commit('setCategories', data.categories || []);
                            this.categories_copy = data.categories;
                            this.update_ok = true;
                            setTimeout(() => {
                                this.update_ok = false;
                            }, 5000);
                        } else {
                            this.update_wrong = true;
                            setTimeout(() => {
                                this.update_wrong = false;
                            }, 5000);
                        }
                        this.editMode = false;
                    })
            },
            cancelChanges() {
                this.editMode = false;
                this.categories_copy = this.$store.categories;
            },
            getUserCategories() {
                if (this.$store.categories) {
                    return;
                }
                this.askBackend('user_category/get', {})
                    .then(({data}) => {
                        if (data.ok) {
                            this.$store.commit('setCategories', data.categories || []);
                            this.categories_copy = data.categories;
                        }
                    })
                    .catch()
            }
        }
    }
</script>

<style scoped lang="scss">
    .cancel_top {
        margin-right: 15px;
    }
    .sub_draft {
        margin-left: 30px;
    }
    .editor_wrapper {
        padding-left: 15px;
        padding-right: 15px;
        padding-bottom: 20px;
        .row {
            margin-left: 15px;
            margin-right: 15px;
        }
    }
    .button_wrapper {
        padding: 20px;
        margin: 10px;
    }
</style>
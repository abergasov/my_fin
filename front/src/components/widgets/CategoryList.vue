<template>
    <v-card class="mx-auto">
        <v-toolbar flat v-bind:class="[(incoming ? 'error' : 'success'), 'cat_wrap']">
            <v-toolbar-title class="font-weight-light">{{ $t((incoming ? 'category' : 'incoming_category')) }}</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn v-if="false && editMode" color="error" class="cancel_top" fab small @click="cancelChanges">
                <v-icon>mdi-close-octagon-outline</v-icon>
            </v-btn>
            <v-btn v-if="editMode" color="success" fab small @click="addRootCategory">
                <v-icon>mdi-plus</v-icon>
            </v-btn>
            <v-btn v-else color="info" fab small @click="startEdit">
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
            <v-list-group v-for="c in categories" no-action sub-group value="true" :key="c.id">
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
</template>

<script>
    export default {
        name: "CategoryList",
        props: {
            incoming: Boolean,
        },
        data () {
            return {
                snackbar: true,
                editMode: false,
                categories_copy: [],
            }
        },
        computed: {
            categories() {
                return this.incoming ? this.$store.state.categories_incoming : this.$store.state.categories_expenses;
            }
        },
        created() {
            this.getUserCategories();
        },
        methods: {
            startEdit() {
                this.categories_copy = this.categories || [];
                this.editMode = !this.editMode
            },
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
            updateCategories() {
                let params = this.incoming ? {cat_in: this.categories_copy} : {cat: this.categories_copy};
                this.askBackend('data/user_category/update', params)
                    .then(data => {
                        this.$store.commit('setAlert', {
                            display: true,
                            text: (data.ok ?  this.$t('categories_updated') : this.$t('categories_updated_error')),
                            color: (data.ok ? 'success' : 'error'),
                        });
                        if (!data.ok) {
                            return;
                        }
                        let type = this.incoming ? 'setCategoriesIncoming' : 'setCategoriesExpenses';
                        this.$store.commit(type, data.categories || []);
                        this.categories_copy = data.categories;
                        this.editMode = false;
                    })
            },
            cancelChanges() {
                this.editMode = false;
                this.categories_copy = this.categories;
            },
            getNewCategory() {
                return {
                    title: '',
                    id: Math.floor(Math.random() * 1000),
                    sub: [],
                };
            },
            getUserCategories() {
                if (this.categories) {
                    this.categories_copy = this.categories;
                }
            }
        }
    }
</script>

<style scoped>
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
    }
    .cat_wrap {

    }
    .button_wrapper {
        padding: 20px;
        margin: 10px;
    }

</style>
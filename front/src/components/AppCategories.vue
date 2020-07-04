<template>
    <v-form width="600">
        <v-container>
            <v-row no-gutters>
                <v-col class="category_container" cols="12" sm="12" md="6">
                    <CategoryList :incoming="true"></CategoryList>
                </v-col>
                <v-col class="category_container" cols="12" sm="12" md="6">
                    <CategoryList :incoming="false"></CategoryList>
                </v-col>
            </v-row>
        </v-container>
    </v-form>
</template>

<script>
    import CategoryList from "./widgets/CategoryList";
    export default {
        components: {CategoryList},
        name: "AppCategories",
        data () {
            return {

            }
        },
        computed: {
            categories() {
                return this.$store.state.categories;
            }
        },
        methods: {

        }
    }
</script>

<style scoped lang="scss">
    .category_container {
        padding: 5px;
    }
</style>
<template>
    <v-card class="mx-auto" outlined>
        <v-toolbar flat color="green darken-2">
            <v-icon>mdi-account</v-icon>
            <v-toolbar-title class="font-weight-light">User Profile</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn color="green darken-4" fab small @click="isEditing = !isEditing">
                <v-icon v-if="isEditing">mdi-close</v-icon>
                <v-icon v-else>mdi-pencil</v-icon>
            </v-btn>
        </v-toolbar>

        <v-list-item three-line>
            <v-list-item-content>
                <v-row>
                    <v-col cols="12" md="4">
                        <v-text-field
                                v-model="firstname"
                                :disabled="!isEditing"
                                :counter="10"
                                label="First name"
                                required
                        ></v-text-field>
                    </v-col>
                    <v-col cols="12" md="4"></v-col>
                    <v-col cols="12" md="4"></v-col>
                </v-row>
            </v-list-item-content>
        </v-list-item>

        <v-card-actions>
            <v-spacer />
            <v-btn color="success" large>{{ $t('update') }}</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
    export default {
        name: "AppProfile",
        data () {
            return {
                isEditing: false,
            }
        },
    }
</script>

<style scoped>

</style>
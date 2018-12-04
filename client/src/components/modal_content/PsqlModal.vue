<template>
    <v-layout row wrap>
        <v-flex xs12>
            <span class="body-2">Database Information</span>
        </v-flex>
        <v-flex xs12 v-for="item in db_information" :key="item.title">
            <v-flex xs4>
                <span>{{ item.title }}</span>
            </v-flex>
            <v-flex xs6>
                <span>: {{ item.value }}</span>
            </v-flex>
        </v-flex>
    </v-layout>
</template>

<script>
export default {
    data(){
        return {
            db_information: [
                { title: "Database", value: "" },
                { title: "Size", value: "" },
            ],
            table_information: [
                { title: "Schema", value: "" },
                { title: "Table", value: "" },
                { title: "Table Size", value: "" },
                { title: "Index Size", value: "" },
            ]
        }
    },
}
</script>

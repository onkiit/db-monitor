<template>
    <v-layout row wrap>
        <v-flex xs12>
            <span class="body-2">Database Information</span>
        </v-flex>
        <v-flex xs12 v-for="item in db_information" :key="item.key">
            <v-layout row wrap>
                <v-flex xs4>
                    <span>{{ item.title }}</span>
                </v-flex>
                <v-flex xs6>
                    <span v-if="item.key==='dbname'">: {{ health.db_information.db_name }}</span>
                    <span v-if="item.key==='dbsize'">: {{ health.db_information.db_size }} KB</span>
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex xs12>
            <span class="body-2">Table Information</span>
        </v-flex>
        <v-flex xs12 v-for="(item, index) in health.table_information" :key="index">
            <v-layout row wrap>
                <v-flex xs4>
                    <span class="body-2">Schema</span>
                </v-flex>
                <v-flex xs6>
                    <span class="body-2">: {{ item.schema_name }}</span>
                </v-flex>
            </v-layout>
           <v-layout row wrap>
                <v-flex xs4>
                    <span> - Table</span>
                </v-flex>
                <v-flex xs6>
                    <span>: {{ item.table_name }}</span>
                </v-flex>
            </v-layout>
            <v-layout row wrap>
                <v-flex xs4>
                    <span> - Table Size</span>
                </v-flex>
                <v-flex xs6>
                    <span>: {{ item.table_size }} KB</span>
                </v-flex>
            </v-layout>
            <v-layout row wrap>
                <v-flex xs4>
                    <span> - Index Size</span>
                </v-flex>
                <v-flex xs6>
                    <span>: {{ item.index_size }} KB</span>
                </v-flex>
            </v-layout>
            <v-layout row wrap class="mb-1">
                <v-flex xs12>
                    <v-divider></v-divider>
                </v-flex>
            </v-layout>
        </v-flex>
    </v-layout>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
    data(){
        return {
            db_information: [
                { title: "Database", key: "dbname" },
                { title: "Size", key: "dbsize" },
            ]
        }
    },
    computed: {
        ...mapGetters({
            health: 'psqlHealth'
        }),
        table_length(){
            return this.health.table_information.length
        }
    }
}
</script>

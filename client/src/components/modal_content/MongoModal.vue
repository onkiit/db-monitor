<template>
    <v-layout row wrap>
        <v-flex xs12 v-for="item in health" :key="item.title">
            <v-layout row wrap>
                <v-flex xs4>
                    <span class="body-2">{{ item.title }}</span>
                </v-flex>
                <v-flex xs6>
                    <span>: {{ item.value }}</span>
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
            health: [
                { title: "Database Name", value: this.$store.getters.mongoHealth.dbname },
                { title: "Available Collection", value: this.$store.getters.mongoHealth.available_collection },
                { title: "Storage Size", value: this.$store.getters.mongoHealth.storage_size },
                { title: "Indexes", value: this.$store.getters.mongoHealth.indexes },
                { title: "Data Size", value: this.$store.getters.mongoHealth.data_size },
            ]
        }
    },
    computed: {
        ...mapGetters({
            mongoHealth: 'mongoHealth'
        })
    }
}
</script>

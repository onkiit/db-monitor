<template>
    <v-layout row wrap>
        <v-flex xs12 v-for="item in mongoHealth" :key="item.title">
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
            mongoHealth: [
                { title: "Database Name", value: this.$store.getters['mongo/health'].dbname },
                { title: "Available Collection", value: this.$store.getters['mongo/health'].available_collection },
                { title: "Storage Size", value: this.$store.getters['mongo/health'].storage_size },
                { title: "Indexes", value: this.$store.getters['mongo/health'].indexes },
                { title: "Data Size", value: this.$store.getters['mongo/health'].data_size },
            ]
        }
    }
}
</script>

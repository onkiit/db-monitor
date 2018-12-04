<template>
    <v-card height="250">
        <v-card-title>
            <span class="title">Mongo</span>
            <v-spacer></v-spacer>
            <v-icon @click="modal()" color="primary">info</v-icon>
        </v-card-title>
        <v-card-text>
            <v-layout row wrap>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Version</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption">{{ monitoringData.version }}</span>
                    </v-flex>
                </v-flex>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Active Clients</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption"><strong>{{ monitoringData.active_client }}</strong> {{monitoringData.active_client > 2 ? "Active clients": "Active client"}}</span>
                    </v-flex>
                </v-flex>
            </v-layout>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    data(){
        return{
            monitoringData: {
                version: "",
                active_client: 0
            }
        }
    },
    created(){
        this.getVersion(),
        this.getActiveClient()
    },
    methods: {
        getVersion(){
            this.$http.get("/mongo/version")
            .then(({ data }) => {
                this.monitoringData.version = data.version
            })
        },
        getActiveClient(){
            this.$http.get("/mongo/client")
            .then(({ data }) => {
                this.monitoringData.active_client = data.active_client
            })
        },
        modal(){
            this.$store.commit('setModal', true)
            this.$store.commit('setModalTitle', 'Mongo Health Information')
            this.$store.commit('setModalCaller', 'mongo')
        }
    }
}
</script>
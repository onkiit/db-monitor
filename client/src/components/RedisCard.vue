<template>
    <v-card height="250">
        <v-card-title>
            <span class="title">Redis</span>
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
                        <span class="caption">{{ monitoringData.active_client }}</span>
                    </v-flex>
                </v-flex>
            </v-layout>
        </v-card-text>
    </v-card>
</template>

<script>
import axios from 'axios'
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
        this.getVersion()
        this.getActiveClient()
    },
    methods: {
        getVersion(){
            axios.get("http://127.0.0.1:8180/redis/version")
            .then(({data}) => {
                this.monitoringData.version = data.version
            })
        },
        getActiveClient(){
            axios.get("http://127.0.0.1:8180/redis/client")
            .then(({data}) => {
                this.monitoringData.active_client = data.active_client
            })
        }
    }
}
</script>
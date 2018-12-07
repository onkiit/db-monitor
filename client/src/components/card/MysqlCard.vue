<template>
    <v-card height="250">
        <v-card-title>
            <span class="title">Mongo</span>
            <v-spacer></v-spacer>
            <v-icon @click="$store.getters['mysql/disabled'] ? '': modal()" :color="$store.getters['mysql/disabled'] ? 'grey lighten-2':'primary'">info</v-icon>
        </v-card-title>
        <v-card-text>
            <v-layout row wrap v-if="$store.getters['mysql/disabled']" align-center>
                <v-flex xs12 justify-center>
                    <span class="body-2">MySQL is disabled by server.</span>
                </v-flex>
            </v-layout>
            <v-layout row wrap v-else>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Version</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption">{{ version }}</span>
                    </v-flex>
                </v-flex>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Active Clients</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption"><strong>{{ active_client }}</strong> {{active_client > 2 ? "Active clients": "Active client"}}</span>
                    </v-flex>
                </v-flex>
            </v-layout>
        </v-card-text>
    </v-card>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'
export default {
    created(){
        this.setVersion
        this.setActiveClient
        this.setHealth
    },
    computed: {
        ...mapGetters({
            version: 'mysql/version',
            active_client: 'mysql/active_client',
            health: 'mysql/health'
        }),
        ...mapActions({
            setVersion: 'mysql/setVersion',
            setActiveClient: 'mysql/setActiveClient',
            setHealth: 'mysql/setHealth'
        })
    },
    methods: {
        ...mapMutations({
            setModal: 'setModal',
            setModalTitle: 'setModalTitle',
            setModalCaller: 'setModalCaller',
        }),
        modal(){
            this.setModal(true)
            this.setModalTitle('Mongo Health Information')
            this.setModalCaller('mysql')
        }
    }
}
</script>
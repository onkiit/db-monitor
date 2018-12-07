<template>
    <v-card height="250">
        <v-card-title>
            <span class="title">PostgreSQL</span>
            <v-spacer></v-spacer>
            <v-icon @click="$store.getters['postgres/disabled'] ? '': modal()" :color="$store.getters['postgres/disabled'] ? 'grey lighten-2':'primary'">info</v-icon>
        </v-card-title>
        <v-card-text>
            <v-layout v-if="$store.getters['postgres/disabled']" row wrap>
                <span class="body-2">PostgreSQL is disabled by server.</span>
            </v-layout>
           <v-layout v-else row wrap>
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
            version: 'postgres/version',
            active_client: 'postgres/active_client',
            health: 'postgres/health'
        }),
        ...mapActions({
            setVersion: 'postgres/setVersion',
            setActiveClient: 'postgres/setActiveClient',
            setHealth: 'postgres/setHealth'
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
            this.setModalTitle('Postgres Health Information')
            this.setModalCaller('postgres')
        }
    }
}
</script>
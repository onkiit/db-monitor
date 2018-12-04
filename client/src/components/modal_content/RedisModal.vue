<template>
    <v-layout row wrap>
        <v-flex xs12 v-for="item in redisHealth" :key="item.title">
            <v-layout row wrap class="mt-1">
                <v-flex xs4>
                    <span class="body-2">{{ item.title }}</span>
                </v-flex>
                <v-flex xs8>
                    <span class="body-1">: {{ item.value }}</span>
                </v-flex>
            </v-layout>
        </v-flex>
    </v-layout>
</template>

<script>
export default {
    data(){
        return {
            redisHealth: [
                { title: "Available Keys", value: this.$store.state.redisHealth.available_key },
                { title: "Expired Keys", value: this.$store.state.redisHealth.expired_key },
                { title: "Evicted Keys", value: this.$store.state.redisHealth.evicted_key },
                { title: "Memory Usage", value: this.$store.state.redisHealth.memory_usage },
                { title: "SlowLog Count", value: this.$store.state.redisHealth.slow_count },
                { title: "Memory Stats", value: "" },
                { title: " - Peak Allocated", value: this.$store.state.redisHealth.memory_stats.peak_allocated },
                { title: " - Total Allowed", value: this.$store.state.redisHealth.memory_stats.total_allowed },
                { title: " - Startup Allocated", value: this.$store.state.redisHealth.memory_stats.startup_allocated },
            ]
        }
    },
    
}
</script>

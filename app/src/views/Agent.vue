<script setup lang="ts">
import type { Host } from '@/service/AgentService';
import { useAgentStore } from '@/stores/agent';

const props = defineProps({
    uuid: {
        type: String,
        required: true
    }
})

const host: Host = useAgentStore().agents[props.uuid]

</script>
<template>
    <div>
        <p-card>
            <template #content>
                <div style="white-space:pre">
                    {{ JSON.stringify(host, null, 2) }}
                </div>
            </template>
        </p-card>
    </div>
</template>

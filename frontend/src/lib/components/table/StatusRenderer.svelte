<script lang="ts">
    import { ConsultStatus } from "$lib/types/dashboard"
    export let status: ConsultStatus;
</script>

<div class="flex items-center justify-center w-4 h-4">
    {#if status === ConsultStatus.red}
        <div class="bg-red-500 rounded w-4 h-4"></div>
    {:else if status === ConsultStatus.yellow}
        <div class="bg-yellow-500 rounded w-4 h-4"></div>
    {:else if status === ConsultStatus.green}
        <div class="bg-green-500 rounded w-4 h-4"></div>
    {/if}
</div>
<script lang="ts">
    export let items: string[];
</script>

<div>
    {#if items.length === 0}
        {''}
    {:else}
        {items.join(', ')}
    {/if}
</div>
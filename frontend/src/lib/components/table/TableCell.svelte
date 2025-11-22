<script lang="ts">
    import { TableRenderer } from "$lib/types/table";
    import ArrayRenderer from "./ArrayRenderer.svelte";
    import StatusRenderer from "./StatusRenderer.svelte";
    export let row: any
    export let renderer: TableRenderer | undefined;
</script>

<div>
    {#if renderer === undefined}
        {row}
    {:else if renderer === TableRenderer.status}
        <StatusRenderer status={row} />
    {:else if renderer === TableRenderer.array}
       <ArrayRenderer items={row} />
    {:else}
        {row}
    {/if}
</div>
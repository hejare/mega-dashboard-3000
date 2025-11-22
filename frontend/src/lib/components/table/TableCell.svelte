<script lang="ts">
    import { TableRenderer } from "$lib/types/table";
    import ArrayRenderer from "./ArrayRenderer.svelte";
    import DateRenderer from "./DateRenderer.svelte";
    import StatusRenderer from "./StatusRenderer.svelte";
    export let data: any
    export let renderer: TableRenderer | undefined;
</script>

<div>
    {#if renderer === undefined}
        {data}
    {:else if renderer === TableRenderer.status}
        <StatusRenderer status={data} />
    {:else if renderer === TableRenderer.array}
       <ArrayRenderer items={data} />
    {:else if renderer === TableRenderer.date}
        <DateRenderer date={data} />
    {:else}
        {data}
    {/if}
</div>
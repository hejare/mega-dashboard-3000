<script lang="ts">
  import TableCell from "$lib/components/table/TableCell.svelte";
  import { ConsultStatus } from "$lib/types/dashboard";
  import { TableRenderer } from "$lib/types/table";
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte"

  export let data

  type Column<T> = {
    key: string,
    displayName: string,
    display?: TableRenderer
  }

  type Row = {
    consultantName: string,
    currentAssignment: string,
    availableFrom: string,
    changeStatus: ConsultStatus,
    extensionStatus: string,
    leads: string[]
  }

  const columns: Column<Row>[] = [{
    key: 'consultantName',
    displayName: 'Konsult',
  }, {
    key: 'currentAssignment',
    displayName: 'Nuvarande uppdrag',
  }, {
    key: 'availableFrom',
    displayName: 'Tillgänglig från',
    display: TableRenderer.date
  }, {
    key: 'changeStatus',
    displayName: 'Status',
    display: TableRenderer.status
  }, {
    key: 'extensionStatus',
    displayName: 'Förlängning',
  }, {
    key: 'leads',
    displayName: 'Leads',
    display: TableRenderer.array
  }];
</script>

<h1 class="text-3xl font-bold underline">Dashboard</h1>
<a href="/" class="text-blue-500 hover:underline">
  Back to Home
</a>

{#if !data}
  <p>Loading...</p>
{:else if data.items.length === 0}
  <p>No data available.</p>
{:else}
  <Table class="mt-4">
      <TableHead>
          {#each columns as col}
              <TableHeadCell>{col.displayName}</TableHeadCell>
          {/each}
      </TableHead>
      <TableBody>
          {#each data.items as row}
              <TableBodyRow>
                  {#each columns as col}
                      <TableBodyCell><TableCell data={row[col.key]} renderer={col.display} /></TableBodyCell>
                  {/each}
              </TableBodyRow>
          {/each}
      </TableBody>
  </Table>
{/if}

<script lang="ts">
    import TableCell from "$lib/components/table/TableCell.svelte";
  import { ConsultStatus } from "$lib/types/dashboard";
  import { TableRenderer } from "$lib/types/table";
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte"

  type Column<T> = {
    key: string,
    displayName: string,
    display?: TableRenderer
  }

  type Row = {
    consultant: string,
    assignment: string,
    availability: string,
    status: ConsultStatus,
    extension: string,
    leads: string[]
  }

  const columns: Column<Row>[] = [{
    key: 'consultant',
    displayName: 'Konsult',
  }, {
    key: 'assignment',
    displayName: 'Nuvarande uppdrag',
  }, {
    key: 'availability',
    displayName: 'Tillgänglig från',
  }, {
    key: 'status',
    displayName: 'Status',
    display: TableRenderer.status
  }, {
    key: 'extension',
    displayName: 'Förlängning',
  }, {
    key: 'leads',
    displayName: 'Leads',
    display: TableRenderer.array
  }];
  const data: Row[] = [{
    consultant: 'Anna Svensson',
    assignment: 'Relex',
    availability: '2024-07-01',
    status: ConsultStatus.red,
    extension: '',
    leads: ['SVT', 'Voi']
  }, {
    consultant: 'Erik Johansson',
    assignment: 'Scrive',
    availability: '2024-08-15',
    status: ConsultStatus.yellow,
    extension: 'Ja',
    leads: []
  }, {
    consultant: 'Maria Karlsson',
    assignment: 'Utveckling',
    availability: '2024-09-01',
    status: ConsultStatus.red,
    extension: '',
    leads: ['SVT']
  }]
</script>

<h1 class="text-3xl font-bold underline">Dashboard</h1>
<a href="/" class="text-blue-500 hover:underline">
  Back to Home
</a>

<Table class="mt-4">
    <TableHead>
        {#each columns as col}
            <TableHeadCell>{col.displayName}</TableHeadCell>
        {/each}
    </TableHead>
    <TableBody>
        {#each data as row}
            <TableBodyRow>
                {#each Object.entries(row) as [k, v]}
                    <TableBodyCell><TableCell row={v} renderer={columns.find(col => col.key === k)?.display} /></TableBodyCell>
                {/each}
            </TableBodyRow>
        {/each}
    </TableBody>
</Table>

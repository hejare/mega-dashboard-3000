<script lang="ts">
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";

  const columns: { 
    displayName: string
  }[] = [{
    displayName: 'Konsult',
  }, {
    displayName: 'Nuvarande uppdrag'
  }, {
    displayName: 'Tillgänglig från'
  }, {
    displayName: 'Status'
  }, {
    displayName: 'Förlängning'
  }, {
    displayName: 'Leads'
  }];
  const data = [{
    konsult: 'Anna Svensson',
    uppdrag: 'Relex',
    tillgänglig: '2024-07-01',
    status: 'Röd',
    förlängning: '',
    leads: ['SVT', 'Voi']
  }, {
    konsult: 'Erik Johansson',
    uppdrag: 'Scrive',
    tillgänglig: '2024-08-15',
    status: 'Gul',
    förlängning: 'Ja',
    leads: []
  }, {
    konsult: 'Maria Karlsson',
    uppdrag: 'Utveckling',
    tillgänglig: '2024-09-01',
    status: 'Röd',
    förlängning: '',
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
                    <TableBodyCell>{v}</TableBodyCell>
                {/each}
            </TableBodyRow>
        {/each}
    </TableBody>
</Table>



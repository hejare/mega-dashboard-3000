import { useQuery } from '@tanstack/react-query';
import {
  createColumnHelper,
  useReactTable,
  getCoreRowModel,
  flexRender,
} from '@tanstack/react-table';
import { getDashboard } from '../queries/queries';

interface Consultant {
    id: number;
    name: string;
    changeStatus: string;
    probableExtensionStatus: string;
}

interface Lead {
    id: number
    organization: string
    stack: string[]
    role: string
    contact: string
    consulatantId: number
}

interface Assignment {
    id: number
    organization: string
    stack: string[]
    role: string
    contact: string
    hourlyPrice: number
    periodStartAt: Date
    periodEndAt: Date
    lead: Lead
    consultant: Consultant
}

interface DashboardData {
    consultant: Consultant;
    leads: Lead[];
    assignment: Assignment;
}

export function Dashboard() {
    let { data, isLoading } = useQuery({ queryKey: ['dashboard'], queryFn: getDashboard })

    const columnHelper = createColumnHelper<DashboardData>();

    const columns = [
        columnHelper.accessor('consultant.name', {
            header: 'Konsult',
            cell: info => info.getValue(),
        }),
        columnHelper.accessor('assignment.organization', {
            header: 'Nuvarande uppdrag',
            cell: info => info.getValue(),
        }),
        columnHelper.accessor('assignment.periodEndAt', {
            header: 'Tillgänglig fr.o.m.',
            cell: info => info.getValue(),
        }),
        columnHelper.accessor('consultant.changeStatus', {
            header: 'Status',
            cell: info => info.getValue(),
        }),
        columnHelper.accessor('consultant.probableExtensionStatus', {
            header: 'Förlängning?',
            cell: info => info.getValue(),
        }),
        columnHelper.accessor('leads', {
            header: 'Leads',
            cell: info => {
                const leads = info.getValue();
                console.log('lead', leads)
                return leads?.map(lead => (
                    <div key={lead.id}>
                        {lead.organization} - {lead.role}
                    </div>
                ));
            },
        }),
    ];

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
    });

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <table className="table" border={1} style={{ borderCollapse: 'collapse', width: '100%' }}>
            <thead>
                {table.getHeaderGroups().map(headerGroup => {
                    return (
                    <tr key={headerGroup.id}>
                        {headerGroup.headers.map(header => ( // map over the headerGroup headers array
                        <th key={header.id} colSpan={header.colSpan}>
                            {<>{header.column.columnDef.header?.valueOf()}</>}
                        </th>
                        ))}
                    </tr>
                    )
                })}
            </thead>
            <tbody>
                {table.getRowModel().rows.map(row => (
                <tr key={row.id}>
                    {row.getVisibleCells().map(cell => (
                    <td key={cell.id} style={{ padding: '4px' }}>
                        {flexRender(
                            cell.column.columnDef.cell,   // <- calls your custom cell renderer
                            cell.getContext()
                        )}
                    </td>
                    ))}
                </tr>
                ))}
            </tbody>
        </table>
    )
}

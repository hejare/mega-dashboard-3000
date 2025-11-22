import {
  createColumnHelper,
  useReactTable,
  getCoreRowModel,
} from '@tanstack/react-table';

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

const dummyData: DashboardData[] = [{
    consultant: {
        id: 1,
        name: 'John Doe',
        changeStatus: 'Stable',
        probableExtensionStatus: 'High',
    },
    leads: [],
    assignment: {
        id: 1,
        organization: 'Tech Corp',
        stack: ['React', 'TypeScript'],
        role: 'Frontend Developer',
        contact: 'john.doe@example.com',
        hourlyPrice: 75,
        periodStartAt: new Date('2023-01-01'),
        periodEndAt: new Date('2023-12-31'),
        lead: {
            id: 1,
            organization: 'Tech Corp',
            stack: ['React', 'TypeScript'],
            role: 'Frontend Developer',
            contact: 'john.doe@example.com',
            consulatantId: 1,
        },
        consultant: {
            id: 1,
            name: 'John Doe',
            changeStatus: 'Stable',
            probableExtensionStatus: 'High',
        },
    },
}];

const columnHelper = createColumnHelper<DashboardData>();

const columns = [
    columnHelper.accessor('consultant.name', {
        header: 'Consultant Name',
        cell: info => info.getValue(),
    }),
    columnHelper.accessor('consultant.changeStatus', {
        header: 'Change Status',
        cell: info => info.getValue(),
    }),
    columnHelper.accessor('consultant.probableExtensionStatus', {
        header: 'Probable Extension Status',
        cell: info => info.getValue(),
    }),
]

export function Dashboard() {
    const table = useReactTable({
        data: dummyData,
        columns,
        getCoreRowModel: getCoreRowModel(),
    });

    return (
        <table border={1} style={{ borderCollapse: 'collapse', width: '100%' }}>
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
                        {<>{cell.getValue()}</>}
                    </td>
                    ))}
                </tr>
                ))}
            </tbody>
        </table>
    )
}

import { useQuery } from '@tanstack/react-query';
import {
  createColumnHelper,
  useReactTable,
  getCoreRowModel,
  flexRender,
} from '@tanstack/react-table';
import { getDashboard } from '../queries/queries';
import { useState } from 'react';

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
            cell: info => {
                const v = info.getValue()
                if (!v) return '';
                return v.toLocaleString('sv-SE', { year: 'numeric', month: '2-digit', day: '2-digit' })
            },
        }),
        columnHelper.accessor('consultant.changeStatus', {
            header: 'Status',
            cell: info => {
                const [v, setV] = useState<string>(info.getValue());
                switch(v) {
                    case 'Ja': {
                        return <div className="h-4 w-4 bg-yellow-500 cursor-pointer" onClick={() => setV('Nej')}></div>
                    };
                    case 'Nej': {
                        return <div className="h-4 w-4 bg-red-500 cursor-pointer" onClick={() => setV('Ja')}></div>
                    }
                    default: {
                        return <></>
                    }
                }
            },
        }),
        columnHelper.accessor('consultant.probableExtensionStatus', {
            header: 'Förlängning?',
            cell: info => {
                const [v, setV] = useState<string>(info.getValue());



                return (
                    <input value={v} onChange={e => setV(e.target.value)}/>
                );
            },
        }),
        columnHelper.accessor('leads', {
            header: 'Leads',
            cell: info => {
                const leads = info.getValue();
                return leads.map(lead => (lead.organization)).join(', ');
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
        <div className="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
            <table className="table">
                <thead>
                    {table.getHeaderGroups().map(headerGroup => {
                        return (
                        <tr key={headerGroup.id}>
                            {headerGroup.headers.map(header => (
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
                                cell.column.columnDef.cell,
                                cell.getContext()
                            )}
                        </td>
                        ))}
                    </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

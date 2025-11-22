import { useMutation, useQuery } from '@tanstack/react-query';
import {
  createColumnHelper,
  useReactTable,
  getCoreRowModel,
  flexRender,
} from '@tanstack/react-table';
import { createAssignMentMutation, getDashboard } from '../queries/queries';
import { useRef, useState } from 'react';

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
    consultantId: number
    title: string
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

enum ModalState {
    lead = 'LEAD',
    createAssignment = 'CREATE_ASSIGNMENT',
}

export function Dashboard() {
    let { data, isLoading, refetch } = useQuery({ queryKey: ['dashboard'], queryFn: getDashboard })
    const createAssignment = useMutation({
        mutationFn: createAssignMentMutation,
        onSuccess: () => {
            refetch();
        }
    })

    const modalRef =  useRef<HTMLDialogElement>(null);
    const [currentLead, setCurrentLead] = useState<Lead | null>(null);
    const [modalState, setModalState] = useState<ModalState>(ModalState.lead);
    const [hourlyPrice, setHourlyPrice] = useState<number>(0);
    const [periodStartAt, setPeriodStartAt] = useState<string>('');
    const [periodEndAt, setPeriodEndAt] = useState<string>('');
    const [currentConsultantId, setCurrentConsultantId] = useState<number | null>(null);

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
                return new Date(v).toLocaleDateString('sv-SE');
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
                return <div className="flex flex-row gap-2">{leads.map(lead => <button className="cursor-pointer border border-base-300 bg-base-100 px-2 py-1" key={lead.id} onClick={() => {
                    if (modalRef.current) {
                        modalRef.current.showModal();
                        setCurrentLead(lead);
                        setModalState(ModalState.lead);
                        setCurrentConsultantId(info.row.original.consultant.id);
                    }
                }}>{lead.organization}</button>)}</div>;
            },
        }),
    ];

    const table = useReactTable({
        data: (data || []).sort((a: DashboardData, b: DashboardData) => {
            if (!a.assignment) return -1;
            if (!b.assignment) return 1;
            return new Date(a.assignment.periodEndAt).getTime() - new Date(b.assignment.periodEndAt).getTime()
        }),
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
            <dialog id="my_modal_2" className="modal" ref={modalRef}>
                {currentLead && (
                    <>
                        {modalState === ModalState.lead ? (
                            <div className="modal-box">
                                <h3 className="font-bold text-lg">Lead Detaljer</h3>
                                <p>Organisation: {currentLead.organization}</p>
                                <p>Roll: {currentLead.role}</p>
                                <p>Kontakt: {currentLead.contact}</p>
                                <p>Stack: {currentLead.stack ? currentLead.stack.join(', ') : ''}</p>
                                    <button className="cursor-pointer border border-base-300 bg-base-100 px-2 py-1" onClick={() => {
                                        setModalState(ModalState.lead);
                                        if (modalRef.current) {
                                            modalRef.current.close();
                                        }
                                    }}>Avbryt</button>
                                    <button className="cursor-pointer border border-base-300 bg-base-100 px-2 py-1" onClick={() => {
                                        setModalState(ModalState.createAssignment);
                                    }}>Konvertera till uppdrag</button>
                            </div>
                        ) :  (
                            <div className="modal-box">
                                <h3 className="font-bold text-lg">Uppdragsdetaljer</h3>
                                <label>Timpris</label>
                                <input type="number" onChange={(e) => setHourlyPrice(parseInt(e.target.value))} />
                                <label>Period start</label>
                                <input type="date" onChange={(e) => setPeriodStartAt(e.target.value)} />
                                <label>Period slut</label>
                                <input type="date" onChange={(e) => setPeriodEndAt(e.target.value)} />
                                    <button className="cursor-pointer border border-base-300 bg-base-100 px-2 py-1" onClick={() => {
                                        if (modalRef.current && currentConsultantId) {
                                            modalRef.current.close();
                                            createAssignment.mutate({
                                                ...currentLead,
                                                hourlyPrice,
                                                periodStartAt: (new Date(periodStartAt)).toISOString(),
                                                periodEndAt: (new Date(periodEndAt)).toISOString(),
                                                leadId: currentLead.id,
                                                consultantId: currentConsultantId,
                                            })
                                        } 
                                    }}>Bekräfta</button>
                            </div>
                        )}
                    </>
                )}
            </dialog>
        </div>
    )
}

import React, { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { search } from '../queries/queries';
import { useDebounce } from '../hooks/useDebounce';

const Search = () => {
    const [searchStr, setSearchStr] = useState('');
    const debouncedSearchStr = useDebounce(searchStr, 500);

    const { data, isLoading } = useQuery({
        queryKey: ['search', debouncedSearchStr],
        queryFn: () => search(debouncedSearchStr),
        enabled: !!debouncedSearchStr,
    });

    if (isLoading) {
        return <div>Loading...</div>;
    }

    const groupedLeads = data?.reduce((acc: any, current: any) => {
        const { id, consultant_name, ...rest } = current;
        if (!acc[id]) {
            acc[id] = { ...rest, id, consultant_names: [] };
        }
        if (consultant_name) {
            acc[id].consultant_names.push(consultant_name);
        }
        return acc;
    }, {});

    const leads = groupedLeads ? Object.values(groupedLeads) : [];

    return (
        <div className="flex flex-col items-center">
            <label className="input mt-10">
            <svg className="h-[1em] opacity-50" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <g
                stroke-linejoin="round"
                stroke-linecap="round"
                stroke-width="2.5"
                fill="none"
                stroke="currentColor"
                >
                <circle cx="11" cy="11" r="8"></circle>
                <path d="m21 21-4.3-4.3"></path>
                </g>
            </svg>
            <input
                type="text"
                placeholder="Sök leads..."
                className="search"
                value={searchStr}
                onChange={(e) => setSearchStr(e.target.value)}
            />
           
            </label>
            {leads.length === 0 && debouncedSearchStr 
            ?  <p>Inga resultat hittades.</p>
            :<ul className="list bg-base-100 rounded-box shadow-md ">
            {leads?.map((lead: any) => (
                <li className="p-4 pb-2 text-lg opacity-60 tracking-wide" key={lead?.id}>
                    {lead?.organization} - {lead?.consultant_names.join(', ')}
                </li>
            ))}
        </ul>
        }
           
            
        </div>
    );
};

export default Search;
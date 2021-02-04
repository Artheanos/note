import { useState } from 'react';

export function useHook(initValue) {
    const [value, setValue] = useState(initValue);
    const onChange = (e) => setValue(e.target.value);

    return [value, setValue, { value, onChange }];
}

/**
 * @param {RequestInfo} input 
 * @param {RequestInit=} init
 */
export async function myFetch(input, init) {
    let response = await fetch(input, init);
    return await response.text();
}
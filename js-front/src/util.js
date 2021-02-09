import { useState } from 'react';

export function useInput(initValue) {
    const [value, setValue] = useState(initValue);
    const onChange = (e) => setValue(e.target.value);

    return [value, setValue, { value, onChange }];
}

/**
 * @param {string} input
 * @param {RequestInit=} init
 */
export async function myFetch(input, init) {
    try {
        let response = await fetch('http://localhost:8090/' + input, { ...init, credentials: 'include' });
        return await response.text();
    } catch (e) {
        alert(e);
    }
}


/**
 * @param {string} input
 * @param {any} body
 * @param {number} expectedStatus
 * @param {RequestInit=} init
 * 
 * @returns {Promise<string>}
 */
export async function myFetchPost(input, body, expectedStatus = 200, init) {
    if (typeof body !== "string") {
        body = JSON.stringify(body);
    }
    console.log(body);
    let response = await fetch('http://localhost:8090/' + input, {
        ...init,
        body,
        method: "POST",
        credentials: 'include',
        headers: {
            "Content-Type": "application/json"
        }
    });
    if (response.status === expectedStatus) {
        return await response.text();
    } else {
        alert("Wrong status");
        throw Error("Wrong status");
    }
}
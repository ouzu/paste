import { base58_to_binary, binary_to_base58 } from 'base58-js'

const keyUsage: KeyUsage[] = [
    "encrypt",
    "decrypt",
];

const algorithm = "AES-CBC";
const keyFormat = "raw";

export function generateIV(): Uint8Array {
    return crypto.getRandomValues(new Uint8Array(16));
}

export function exportIV(iv: Uint8Array): string {
    return binary_to_base58(iv);
}

export function importIV(s: string): Uint8Array {
    return base58_to_binary(s);
}

export async function generateKey(): Promise<CryptoKey> {
    return await crypto.subtle
        .generateKey(
            { name: algorithm, length: 256 },
            true,
            keyUsage,
        );
}

export async function exportKey(key: CryptoKey): Promise<string> {
    let k = await crypto.subtle.exportKey(keyFormat, key);

    return binary_to_base58(new Uint8Array(k));
}

export async function importKey(s: string): Promise<CryptoKey> {
    return await crypto.subtle.importKey(
        keyFormat,
        base58_to_binary(s),
        algorithm,
        true,
        keyUsage
    )
}

export async function encrypt(data: ArrayBuffer, iv: Uint8Array, key: CryptoKey): Promise<ArrayBuffer> {
    return await crypto.subtle.encrypt(
        { name: "AES-CBC", iv },
        key,
        data
    )
}

export async function decrypt(data: ArrayBuffer, iv: Uint8Array, key: CryptoKey): Promise<ArrayBuffer> {
    return await crypto.subtle.decrypt(
        { name: "AES-CBC", iv },
        key,
        data
    )
}


const encoder = new TextEncoder();
export async function encryptString(data: string, iv: Uint8Array, key: CryptoKey): Promise<string> {
    return binary_to_base58(new Uint8Array(await encrypt(encoder.encode(data), iv, key)));
}

const decoder = new TextDecoder();
export async function decryptString(data: string, iv: Uint8Array, key: CryptoKey): Promise<string> {
    return decoder.decode(await decrypt(base58_to_binary(data), iv, key))
}
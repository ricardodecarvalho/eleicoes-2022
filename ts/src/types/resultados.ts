export type ResultadosType = {
    abr: AbrType[]
}

export type AbrType = {
    dt: string
    ht: string
    pst: string
    cand: CandidatosType[]
}

export type CandidatosType = {
    seq: string
    n: string
    vap: string
    pvap: string
    e: string
}

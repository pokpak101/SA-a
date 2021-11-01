import { AssessInterface } from "./IAssess";
import { SymptomInterface } from "./ISymptom";
import { StateInterface } from "./IState";
import { CaseInterface } from "./ICase";

export interface AssessmentsheetInterface{
    ID: number,
    Name:   string,
    AssessTime: Date,
    AssessID:   number,
    Assess:     AssessInterface,
    SymptomID:  number,
    Symptom:    SymptomInterface,
    StateID:    number,
    State:      StateInterface,
    CaseID:     number,
    Case:       CaseInterface,
}
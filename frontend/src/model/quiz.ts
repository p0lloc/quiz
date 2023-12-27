export interface Quiz {
    id: string;
    name: string;
    questions: QuizQuestion[];
}

export interface QuizQuestion {
    id: string;
    name: string;
    choices: QuizChoice[];
}

export interface QuizChoice {
    id: string;
    value: string;
    correct: boolean;
}
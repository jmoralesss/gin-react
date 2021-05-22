import { History, LocationState } from "history";

export type FormEntryState = {
	email: string;
	password: string;
}

export type FormEntryProps = {
	history: History<LocationState>;
}
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';
import {time} from '../models';

export function CreateActivity(arg1:model.Activity):Promise<model.Activity>;

export function CreateOrMergeActivity(arg1:model.Activity):Promise<model.Activity>;

export function DeleteActivity(arg1:string):Promise<void>;

export function GetActivity(arg1:string):Promise<model.Activity>;

export function GetMostRecentDate():Promise<time.Time>;

export function UpdateActivity(arg1:model.Activity):Promise<model.Activity>;

// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';
import {time} from '../models';
import {oauth2} from '../models';

export function BackfillData(arg1:model.User):Promise<void>;

export function GetNewData(arg1:model.User,arg2:time.Time):Promise<void>;

export function OpenStravaAuth():Promise<void>;

export function StartListener():Promise<oauth2.Token>;

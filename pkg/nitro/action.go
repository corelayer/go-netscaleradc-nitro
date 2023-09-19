/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package nitro

var (
	ActionNone    = Action{""}
	ActionCreate  = Action{"create"}
	ActionRename  = Action{"rename"}
	ActionEnable  = Action{"enable"}
	ActionDisable = Action{"disable"}
	ActionCount   = Action{"count"}
	ActionRestore = Action{"restore"}
	ActionSave    = Action{"save"}
	ActionSync    = Action{"sync"}
	ActionForce   = Action{"force"}
	ActionClear   = Action{"clear"}
	ActionLink    = Action{"link"}
	ActionUnlink  = Action{"unlink"}
	ActionUpdate  = Action{"update"}
)

type Action struct {
	string
}

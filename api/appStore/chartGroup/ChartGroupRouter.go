/*
 * Copyright (c) 2020-2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chartGroup

import (
	"github.com/gorilla/mux"
)

type ChartGroupRouterImpl struct {
	ChartGroupRestHandler ChartGroupRestHandler
}
type ChartGroupRouter interface {
	InitChartGroupRouter(helmRouter *mux.Router)
}

func NewChartGroupRouterImpl(ChartGroupRestHandler ChartGroupRestHandler) *ChartGroupRouterImpl {
	return &ChartGroupRouterImpl{ChartGroupRestHandler: ChartGroupRestHandler}

}

func (impl *ChartGroupRouterImpl) InitChartGroupRouter(chartGroupRouter *mux.Router) {
	chartGroupRouter.Path("/").
		HandlerFunc(impl.ChartGroupRestHandler.CreateChartGroup).Methods("POST")
	chartGroupRouter.Path("/").
		HandlerFunc(impl.ChartGroupRestHandler.UpdateChartGroup).Methods("PUT")
	chartGroupRouter.Path("/entries").
		HandlerFunc(impl.ChartGroupRestHandler.SaveChartGroupEntries).Methods("PUT")
	chartGroupRouter.Path("/list").
		HandlerFunc(impl.ChartGroupRestHandler.GetChartGroupList).Methods("GET")
	chartGroupRouter.Path("/{chartGroupId}").
		HandlerFunc(impl.ChartGroupRestHandler.GetChartGroupWithChartMetaData).Methods("GET")
	chartGroupRouter.Path("/installation-detail/{chartGroupId}").
		HandlerFunc(impl.ChartGroupRestHandler.GetChartGroupInstallationDetail).Methods("GET")

	chartGroupRouter.Path("/list/min").
		HandlerFunc(impl.ChartGroupRestHandler.GetChartGroupListMin).Methods("GET")
	chartGroupRouter.Path("").
		HandlerFunc(impl.ChartGroupRestHandler.DeleteChartGroup).Methods("DELETE")
}

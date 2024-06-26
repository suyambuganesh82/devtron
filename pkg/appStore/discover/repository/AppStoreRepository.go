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

package appStoreDiscoverRepository

import (
	dockerArtifactStoreRegistry "github.com/devtron-labs/devtron/internal/sql/repository/dockerRegistry"
	chartRepoRepository "github.com/devtron-labs/devtron/pkg/chartRepo/repository"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"time"
)

type AppStoreRepository interface{}

type AppStoreRepositoryImpl struct {
	dbConnection *pg.DB
	Logger       *zap.SugaredLogger
}

func NewAppStoreRepositoryImpl(Logger *zap.SugaredLogger, dbConnection *pg.DB) *AppStoreRepositoryImpl {
	return &AppStoreRepositoryImpl{dbConnection: dbConnection, Logger: Logger}
}

type AppStore struct {
	TableName   struct{} `sql:"app_store" pg:",discard_unknown_columns"`
	Id          int      `sql:"id,pk"`
	Name        string   `sql:"name,notnull"`
	ChartRepoId int      `sql:"chart_repo_id"`
	//Active                bool      `sql:"active,notnull"`
	DockerArtifactStoreId string    `sql:"docker_artifact_store_id"`
	ChartGitLocation      string    `sql:"chart_git_location"`
	CreatedOn             time.Time `sql:"created_on,notnull"`
	UpdatedOn             time.Time `sql:"updated_on,notnull"`
	ChartRepo             *chartRepoRepository.ChartRepo
	DockerArtifactStore   *dockerArtifactStoreRegistry.DockerArtifactStore
}

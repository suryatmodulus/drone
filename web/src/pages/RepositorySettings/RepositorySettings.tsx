/*
 * Copyright 2023 Harness, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React from 'react'

import { PageBody, Container, Tabs } from '@harnessio/uicore'
import { useGetRepositoryMetadata } from 'hooks/useGetRepositoryMetadata'
import { useStrings } from 'framework/strings'

import { RepositoryPageHeader } from 'components/RepositoryPageHeader/RepositoryPageHeader'
import { getErrorMessage, voidFn } from 'utils/Utils'
import { LoadingSpinner } from 'components/LoadingSpinner/LoadingSpinner'
import GeneralSettingsContent from './GeneralSettingsContent/GeneralSettingsContent'
import css from './RepositorySettings.module.scss'

enum SettingsTab {
  webhooks = 'webhooks',
  general = 'general'
}
export default function RepositorySettings() {
  const { repoMetadata, error, loading, refetch } = useGetRepositoryMetadata()

  const [activeTab, setActiveTab] = React.useState<string>(SettingsTab.general)

  const { getString } = useStrings()
  return (
    <Container className={css.main}>
      <RepositoryPageHeader
        repoMetadata={repoMetadata}
        title={getString('settings')}
        dataTooltipId="repositorySettings"
      />
      <PageBody error={getErrorMessage(error)} retryOnError={voidFn(refetch)}>
        <LoadingSpinner visible={loading} />
        {repoMetadata && (
          <Container className={css.main} padding={'large'}>
            <Tabs
              id="SettingsTabs"
              vertical
              large={false}
              defaultSelectedTabId={activeTab}
              animate={false}
              onChange={(id: string) => setActiveTab(id)}
              tabList={[
                {
                  id: SettingsTab.general,
                  title: getString('general'),
                  panel: <GeneralSettingsContent repoMetadata={repoMetadata} refetch={refetch} />,
                  iconProps: { name: 'cog' }
                }
              ]}></Tabs>
          </Container>
        )}
      </PageBody>
    </Container>
  )
}

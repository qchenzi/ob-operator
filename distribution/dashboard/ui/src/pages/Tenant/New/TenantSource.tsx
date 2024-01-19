import { getAllTenants } from '@/services/tenant';
import { useRequest } from 'ahooks';
import {
  Card,
  Col,
  DatePicker,
  Form,
  Input,
  Row,
  Select,
  Switch,
  TimePicker,
} from 'antd';
import { FormInstance } from 'antd/lib/form';
import { useEffect, useState } from 'react';
import styles from './index.less';

interface TenantSourceProps {
  form: FormInstance<any>;
  clusterName: string;
}
type RoleType = 'PRIMARY' | 'STANDBY';
const { Password } = Input;
export default function TenantSource({ form, clusterName }: TenantSourceProps) {
  const [isChecked, setIsChecked] = useState<boolean>(false);
  const [recoverChecked, setRecoverChecked] = useState<boolean>(false);
  const [synchronizeChecked, setSynchronizeChecked] = useState<boolean>(false);
  const [recoverTimeChecked, setRecoverTimeChecked] = useState<boolean>(false);
  const [selectRole, setSelectRole] = useState<RoleType | undefined>();
  const subTitleStyle = {
    display: 'flex',
    alignItems: 'center',
    marginBottom: 24,
  };
  const h3Style = { marginBottom: 0, marginRight: 12 };
  //主租户和备租户
  const tenantRole = [
    {
      label: 'PRIMARY',
      value: 'PRIMARY',
    },
    {
      label: 'STANDBY',
      value: 'STANDBY',
    },
  ];
  const recoverType = [
    {
      label: 'NFS',
      value: 'NFS',
    },
    {
      label: 'OSS',
      value: 'OSS',
    },
  ];

  const tenantRoleChange = (value: RoleType) => {
    setSelectRole(value);
    if (value === 'PRIMARY' && synchronizeChecked) {
      setSynchronizeChecked(false);
    }
  };

  const { run: getTenants, data: tenantListRes } = useRequest(getAllTenants);

  const tenantList = tenantListRes?.data.map((tenant) => ({
    label: tenant.name,
    value: tenant.name,
  }));
  console.log('tenantList', tenantList);

  const onChange = (date, dateString) => {
    console.log(date, dateString);
  };
  const onChangea = (time: Dayjs, timeString: string) => {
    console.log(time, timeString);
  };
  useEffect(() => {
    if (synchronizeChecked && clusterName) {
      getTenants(clusterName);
    }
  }, [synchronizeChecked, clusterName]);

  return (
    <Card
      title={
        <div>
          标题
          {/* <span>租户资源</span>{' '}
          <Switch onChange={(val) => setIsChecked(val)} checked={isChecked} /> */}
        </div>
      }
    >
      <div>
        <div style={subTitleStyle}>
          <h3 style={h3Style}>租户资源</h3>{' '}
          <Switch onChange={(val) => setIsChecked(val)} checked={isChecked} />
        </div>
        {isChecked && (
          <Form.Item
            name={['tenantRole']}
            label="租户角色"
            style={{ width: '50%' }}
            rules={[
              {
                required: true,
                message: '请选择租户角色',
              },
            ]}
          >
            <Select
              placeholder="请选择"
              onChange={(value) => tenantRoleChange(value)}
              options={tenantRole}
            />
          </Form.Item>
        )}
      </div>
      <div style={{ marginBottom: 24 }}>
        <div style={subTitleStyle}>
          <h3 style={h3Style}>从备份恢复</h3>
          <Switch
            onChange={(val) => setRecoverChecked(val)}
            checked={recoverChecked}
          />
        </div>

        {recoverChecked && (
          <div>
            <Row gutter={[16, 32]}>
              <Col span={8}>
                <Form.Item
                  name={['source', 'restore', 'type']}
                  label="恢复类型"
                  rules={[
                    {
                      required: true,
                      message: '请选择恢复类型',
                    },
                  ]}
                >
                  <Select placeholder="请选择" options={recoverType} />
                </Form.Item>
              </Col>
              <Col span={8}>
                <Form.Item
                  label="OSS AccessID"
                  name={['source', 'restore', 'ossAccessId']}
                  rules={[
                    {
                      required: true,
                      message: '请输入 OSS AccessID',
                    },
                  ]}
                >
                  <Password placeholder="请输入" />
                </Form.Item>
              </Col>
              <Col span={8}>
                <Form.Item
                  label="OSS AccessKey"
                  name={['source', 'restore', 'ossAccessKey']}
                  rules={[
                    {
                      required: true,
                      message: '请输入 OSS AccessKey',
                    },
                  ]}
                >
                  <Password placeholder="请输入" />
                </Form.Item>
              </Col>
              <Col span={8}>
                <Form.Item
                  label="日志归档路径"
                  name={['source', 'restore', 'archiveSource']}
                  rules={[
                    {
                      required: true,
                      message: '请输入日志归档路径',
                    },
                  ]}
                >
                  <Input placeholder="请输入" />
                </Form.Item>
              </Col>
              <Col span={8}>
                <Form.Item
                  label="数据备份路径"
                  name={['source', 'restore', 'bakDataSource']}
                  rules={[
                    {
                      required: true,
                      message: '请输入数据备份路径',
                    },
                  ]}
                >
                  <Input placeholder="请输入" />
                </Form.Item>
              </Col>
              <Col span={8}>
                <Form.Item
                  label="加密密码"
                  name={['source', 'restore', 'bakEncryptionPassword']}
                >
                  <Password placeholder="请输入" />
                </Form.Item>
              </Col>
              <Col span={8}>
                {/* <PasswordInput
              value={passwordVal}
              onChange={setPasswordVal}
              form={form}
              name="rootPassword"
            /> */}
              </Col>
            </Row>
            <div>
              {' '}
              <div style={subTitleStyle}>
                <h4 style={h3Style}>恢复至特定时间</h4>{' '}
                <Switch
                  onChange={(val) => setRecoverTimeChecked(val)}
                  checked={recoverTimeChecked}
                  disabled={synchronizeChecked === true}
                />
              </div>
              {recoverTimeChecked && (
                <Row gutter={[16, 32]}>
                  <Col span={12}>
                    <Form.Item
                      className={styles.dateContainer}
                      label="恢复日期"
                      style={{ width: '100%' }}
                      rules={[
                        {
                          required: true,
                          message: '请选择恢复日期',
                        },
                      ]}
                      name={['source', 'restore', 'until', 'date']}
                    >
                      <DatePicker onChange={onChange} />
                    </Form.Item>
                  </Col>
                  <Col span={12}>
                    <Form.Item
                      className={styles.dateContainer}
                      style={{ width: '100%' }}
                      rules={[
                        {
                          required: true,
                          message: '请选择恢复时间',
                        },
                      ]}
                      label="时分秒"
                      name={['source', 'restore', 'until', 'time']}
                    >
                      <TimePicker onChange={onChangea} />
                    </Form.Item>
                  </Col>
                </Row>
              )}
              {/* <InputNumber /> */}
            </div>
          </div>
        )}
      </div>
      <div>
        <div style={subTitleStyle}>
          <h3 style={h3Style}>同步主租户</h3>
          <Switch
            onChange={(val) => setSynchronizeChecked(val)}
            checked={synchronizeChecked}
            disabled={recoverTimeChecked === true || selectRole === 'PRIMARY'}
          />
        </div>
        {synchronizeChecked && (
          <Form.Item
            style={{ width: '50%' }}
            label="主租户"
            name={['source', 'tenant']}
            rules={[
              {
                required: true,
                message: '请选择',
              },
            ]}
          >
            <Select options={tenantList} placeholder="请选择" />
          </Form.Item>
        )}
      </div>
    </Card>
  );
}

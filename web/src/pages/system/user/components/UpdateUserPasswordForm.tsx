import React, {useEffect} from 'react';
import {Form, Input, Modal, Select} from 'antd';
import { UserListItem} from '../data.d';

export interface UpdateFormProps {
  onCancel: () => void;
  onSubmit: (values: Partial<UserListItem>) => void;
  updatePasswordModalVisible: boolean;
  currentData: Partial<UserListItem>;
}
const FormItem = Form.Item;

const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const UpdateUserPasswordForm: React.FC<UpdateFormProps> = (props) => {
  const [form] = Form.useForm();
  const { Option } = Select;

  const {
    onSubmit,
    onCancel,
    updatePasswordModalVisible,
    currentData,
  } = props;

  useEffect(() => {
    if (form && !updatePasswordModalVisible) {
      form.resetFields();
    }
  }, [props.updatePasswordModalVisible]);

  useEffect(() => {
    if (currentData) {
      form.setFieldsValue({
        ...currentData,
        // deptId:currentData.deptId+'',
      });
    }
  }, [props.currentData]);

  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const handleFinish = (values: { [key: string]: any }) => {
    if (onSubmit) {
      onSubmit(values);
    }
  };

  const renderContent = () => {
    return (
      <>
        <FormItem
          name="id"
          label="主键"
          hidden
        >
          <Input id="update-id" placeholder="请输入主键" />
        </FormItem>
        <FormItem
          name="realName"
          label="用户名"
        >
          <Input id="update-name" placeholder={'请输入用户名'}/>
        </FormItem>
        <FormItem
          name="mobile"
          label="手机号"
        >
          <Input id="update-mobile" placeholder={'请输入手机号'}/>
        </FormItem>
        <FormItem
          name="remark"
          label="备注"
        >
          <Input id="update-remark" placeholder={'请输入备注'}/>
        </FormItem>
        <FormItem
          name="status"
          label="状态"
        >
          <Select id="statusID" placeholder={'请选择状态'}>
            <Option value={0}>禁用</Option>
            <Option value={1}>启用</Option>
          </Select>
        </FormItem>
      </>
    );
  };


  const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel };

  return (
    <Modal
      forceRender
      destroyOnClose
      title="修改用户"
      visible={updatePasswordModalVisible}
      {...modalFooter}
    >
      <Form
        {...formLayout}
        form={form}
        onFinish={handleFinish}
      >
        {renderContent()}
      </Form>
    </Modal>
  );
};

export default UpdateUserPasswordForm;

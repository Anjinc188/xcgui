
enum  window_state_flag_
{
	window_state_flag_nothing  =0x0000,  ///<��
	window_state_whole_leave   =0x0001,  ///<��������
	window_state_body_leave    =0x0002,  ///<����-body
	window_state_top_leave     =0x0004,  ///<����-top
	window_state_bottom_leave  =0x0008,  ///<����-bottom
	window_state_left_leave    =0x0010,  ///<����-left
	window_state_right_leave   =0x0020,  ///<����-right
};

enum  element_state_flag_  //���״̬
{
	element_state_flag_nothing   =window_state_flag_nothing,  ///<��
	element_state_flag_enable    =0x0001,  ///<����
	element_state_flag_disable   =0x0002,  ///<����
	element_state_flag_focus     =0x0004,  ///<����
	element_state_flag_focus_no  =0x0008,  ///<�޽���

	element_state_flag_leave     =0x0010,  ///<����뿪
	element_state_flag_stay      =0x0020,  ///<Ϊ��չģ�鱣��
	element_state_flag_down      =0x0040,  ///<Ϊ��չģ�鱣��
};

enum  button_state_flag_  //���״̬
{
	button_state_flag_leave     =element_state_flag_leave,  ///<����뿪
	button_state_flag_stay      =element_state_flag_stay,   ///<���ͣ��
	button_state_flag_down      =element_state_flag_down,   ///<��갴��

	button_state_flag_check     =0x0080, ///<ѡ��
	button_state_flag_check_no  =0x0100, ///<δѡ��
};

enum comboBox_state_flag_
{
	comboBox_state_flag_leave   =element_state_flag_leave, ///<����뿪
	comboBox_state_flag_stay    =element_state_flag_stay,  ///<���ͣ��
	comboBox_state_flag_down    =element_state_flag_down,  ///<��갴��
};

enum listBox_state_flag_
{
	listBox_state_flag_item_leave  =0x0080, ///<������뿪
	listBox_state_flag_item_stay   =0x0100, ///<�����ͣ��
	listBox_state_flag_item_select =0x0200, ///<��ѡ��
};

enum list_state_flag_
{
	list_state_flag_item_leave  =0x0080, ///<������뿪
	list_state_flag_item_stay   =0x0100, ///<�����ͣ��
	list_state_flag_item_select =0x0200, ///<��ѡ��
};

enum listView_state_flag_
{
	listView_state_flag_item_leave  =0x0080, ///<������뿪
	listView_state_flag_item_stay   =0x0100, ///<�����ͣ��
	listView_state_flag_item_select =0x0200, ///<��ѡ��
};

enum tree_state_flag_
{
	tree_state_flag_item_leave  =0x0080, ///<������뿪
	tree_state_flag_item_select =0x0100, ///<��ѡ��
};

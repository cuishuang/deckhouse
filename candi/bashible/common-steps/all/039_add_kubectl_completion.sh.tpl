sed -i 's/\# \"\\e\[5~\": history-search-backward/\"\\e\[5~\": history-search-backward/' /etc/inputrc
sed -i 's/^\# \"\\e\[6~\": history-search-forward/\"\\e\[6~\": history-search-forward/' /etc/inputrc

sed -i 's/\#force_color_prompt=yes/force_color_prompt=yes/' /root/.bashrc
sed -i 's/01;32m/01;31m/' /root/.bashrc

mkdir -p /etc/bash_completion.d
kubectl completion bash >/etc/bash_completion.d/kubectl

completion="if [ -f /etc/bash_completion ] && ! shopt -oq posix; then . /etc/bash_completion ; fi"
grep -qF -- "$completion" /root/.bashrc || echo "$completion" >> /root/.bashrc

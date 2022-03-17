namespace EventDemo
{
    public class Account
    {
        public event EventHandler<string>? RaiseTransactionApprovedEvent;
        public event EventHandler<OverdraftEventArgs>? OverDraftEvent;

        public string? AccountName { get; set; }
        public decimal Balance { get; set; }

        List<string> _transactions = new List<string>();

        public IReadOnlyList<string> Transactions { get => _transactions.AsReadOnly(); }

        public bool AddDeposit(string depositName, decimal amount)
        {
            _transactions.Add($"Deposited {amount:C2} for {depositName}");
            Balance += amount;
            RaiseTransactionApprovedEvent?.Invoke(this, depositName);
            return true;
        }

        public bool MakePayment(string paymentName, decimal amount, Account? backupAccount = null)
        {

            if (Balance >= amount)
            {
                _transactions.Add($"$Withdrew {amount:C2} for {paymentName}");
                Balance -= amount;
                return true;
            }
            else
            {
                if (backupAccount != null)
                {
                    if (backupAccount.Balance + Balance > amount)
                    {
                        decimal amountNeeded = amount - Balance;
                        OverdraftEventArgs args = new OverdraftEventArgs(amountNeeded, "Extra Info");
                        OverDraftEvent?.Invoke(this, args);
                        if (args.CancelTransaction) return false;
                        bool overDraftSucceded = backupAccount.MakePayment("Overdraft Protection", amountNeeded);
                        if (!overDraftSucceded) return false;
                        AddDeposit("Overdraft Protection Deposit", amountNeeded);
                        _transactions.Add($"Withdrew {amount:C2} for {paymentName}");
                        Balance -= amount;
                        RaiseTransactionApprovedEvent?.Invoke(this, paymentName);
                        
                        return true;
                    }
                    else return false;
                }
                else return false;
            }
        }

    }
}